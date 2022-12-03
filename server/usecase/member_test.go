package usecase_test

import (
	"fmt"
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/server/constants"
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/server/domain/service"
	repo "github.com/RodrigoMattosoSilveira/go-swiss-pairing/server/interface/persistence/memory"
	uc "github.com/RodrigoMattosoSilveira/go-swiss-pairing/server/usecase"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/status"
)

var _ = Describe("Member", func() {
	var useCase uc.IMemberUsecase

	var first string
	var last string
	var email string
	var password string
	var cell string

	BeforeEach(func() {
		repository := repo.NewMemberRepository()
		svc := service.NewMemberService(repository)
		useCase = uc.NewMemberUsecase(repository, svc)
		useCase.Empty()

		first = "Adeline"
		last = "Hodge"
		email = "Adeline.Hodge@yahoo.com"
		password = "oPT14J30I#y4"
		cell = "801 277 6891"
	})
	Describe("Validate the password validator", func() {
		It("works for a valid password", func() {
			Expect(uc.ValidatePassword("oPT14J30I#y4")).To(BeTrue())
		})
	})
	Describe("Validate the create function", func() {
		It("Fails when first name is not provided", func() {
			_, err := useCase.Create("", last, email, password, cell)
			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(Equal("rpc error: code = InvalidArgument desc = first name not provided"))
		})
		It("Fails when last name is not provided", func() {
			_, err := useCase.Create(first, "", email, password, cell)
			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(Equal("rpc error: code = InvalidArgument desc = last name not provided"))
		})
		It("Fails when email is not provided", func() {
			_, err := useCase.Create(first, last, "", password, cell)
			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(Equal("rpc error: code = InvalidArgument desc = email not provided"))
		})
		It("Fails when email is invalid", func() {
			_, err := useCase.Create(first, last, "Adeline.Hodgeyahoo.com", password, cell)
			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(Equal("rpc error: code = InvalidArgument desc = invalid email address"))
		})
		It("Fails when password is not provided", func() {
			_, err := useCase.Create(first, last, email, "", cell)
			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(Equal("rpc error: code = InvalidArgument desc = password not provided"))
		})
		It("Fails when password is invalid", func() {
			_, err := useCase.Create(first, last, email, "oPT14J30Iy4", cell)
			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(Equal("rpc error: code = InvalidArgument desc = invalid password"))
		})
		It("Fails when cell phone number is not provided", func() {
			_, err := useCase.Create(first, last, email, password, "")
			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(Equal("rpc error: code = InvalidArgument desc = cell phone number not provided"))
		})
		It("Fails when cell phone number is invalid", func() {
			_, err := useCase.Create(first, last, email, password, "801 277-689")
			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(Equal("rpc error: code = InvalidArgument desc = invalid cell phone number"))
		})
		It("Works when first, last, email, password, and cell phone number are provided and valid", func() {
			member, _ := useCase.Create(first, last, email, password, cell)
			Expect(member).To(Not(BeNil()))
			Expect(member.First()).To(Equal(first))
			Expect(member.Last()).To(Equal(last))
			Expect(member.Email()).To(Equal(email))
			Expect(member.Password()).To(Equal(password))
			Expect(member.Cell()).To(Equal(cell))
		})
	})
	Describe("Validate the ReadId function", func() {
		It("Works when the id is valid", func() {
			member, _ := useCase.Create(first, last, email, password, cell)
			var id = member.Id()
			memberRead, _ := useCase.ReadById(id)
			Expect(memberRead.Id()).To(Equal(id))
			Expect(memberRead).To(Not(BeNil()))
			Expect(memberRead.First()).To(Equal(first))
			Expect(memberRead.Last()).To(Equal(last))
			Expect(memberRead.Email()).To(Equal(email))
			Expect(memberRead.Password()).To(Equal(password))
			Expect(memberRead.Cell()).To(Equal(cell))
		})
		It("Fails when the id is valid", func() {
			useCase.Create(first, last, email, password, cell)
			invalidId := "invalid-id"
			memberRead, err := useCase.ReadById(invalidId)
			Expect(memberRead).To(BeNil())
			Expect(err).To(Not(BeNil()))
			msg := status.Error(constants.GRPC_STATUS_NOT_FOUND, fmt.Sprintf("did not find member with id: %s", invalidId)).Error()
			Expect(err.Error()).To(Equal(msg))
		})
	})
	Describe("Validate the ReadEmail function", func() {
		It("Works when the email is valid", func() {
			member, _ := useCase.Create(first, last, email, password, cell)
			var id = member.Id()
			memberRead, _ := useCase.ReadByEmail(email)
			Expect(memberRead.Id()).To(Equal(id))
			Expect(memberRead).To(Not(BeNil()))
			Expect(memberRead.First()).To(Equal(first))
			Expect(memberRead.Last()).To(Equal(last))
			Expect(memberRead.Email()).To(Equal(email))
			Expect(memberRead.Password()).To(Equal(password))
			Expect(memberRead.Cell()).To(Equal(cell))
		})
		It("Fails when the email is invalid", func() {
			useCase.Create(first, last, email, password, cell)
			invalidEmail := "mario-new@yahoo.com"
			memberRead, err := useCase.ReadByEmail(invalidEmail)
			Expect(memberRead).To(BeNil())
			Expect(err).To(Not(BeNil()))
			msg := status.Error(constants.GRPC_STATUS_NOT_FOUND, fmt.Sprintf("did not find member with email: %s", invalidEmail)).Error()
			Expect(err.Error()).To(Equal(msg))
		})
	})
	Describe("Validate the Read function", func() {
		It("Works when the db is empty", func() {
			members, err := useCase.Read()
			Expect(members).To(Not(BeNil()))
			Expect(err).To(BeNil())
			Expect(len(members)).To(Equal(0))
		})
		It("Works when the db contains one element", func() {
			var first = "mario"
			var email = "mario@yahoo.com"
			// create one member
			member, error := useCase.Create(first, last, email, password, cell)
			Expect(member).To(Not(BeNil()))
			Expect(error).To(BeNil())
			// read all members
			members, err := useCase.Read()
			Expect(members).To(Not(BeNil()))
			Expect(err).To(BeNil())
			Expect(len(members)).To(Equal(1))
			member = members[0]
			Expect(member.First()).To(Equal(first))
			Expect(member.Email()).To(Equal(email))
		})
		It("Works when the db contains multiple elements", func() {
			// create two members
			useCase.Create(first, last, email, password, cell)

			var first_2 = "Maryjane"
			var last_2 = "Lowe"
			var email_2 = "Maryjane.Lowe@yahoo.com"
			var password_2 = "9gv24yPGeEEZ#y4"
			var cell_2 = "801 277-6892"
			useCase.Create(first_2, last_2, email_2, password_2, cell_2)

			// read all members
			members, err := useCase.Read()
			Expect(members).To(Not(BeNil()))
			Expect(err).To(BeNil())
			Expect(len(members)).To(Equal(2))

			member := members[0]
			Expect(member.First()).To(Equal(first))
			Expect(member.Last()).To(Equal(last))
			Expect(member.Email()).To(Equal(email))
			Expect(member.Password()).To(Equal(password))
			Expect(member.Cell()).To(Equal(cell))
			member = members[1]
			Expect(member.First()).To(Equal(first_2))
			Expect(member.Last()).To(Equal(last_2))
			Expect(member.Email()).To(Equal(email_2))
			Expect(member.Password()).To(Equal(password_2))
			Expect(member.Cell()).To(Equal(cell_2))
		})
	})
})
