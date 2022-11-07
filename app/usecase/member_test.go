package usecase_test

import (
	"fmt"
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/constants"
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/domain/service"
	repo "github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/interface/persistence/memory"
	uc "github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/usecase"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/status"
)

var _ = Describe("Member", func() {
	var useCase uc.IMemberUsecase
	BeforeEach(func() {
		repository := repo.NewMemberRepository()
		svc := service.NewMemberService(repository)
		useCase = uc.NewMemberUsecase(repository, svc)
		useCase.Empty()
	})
	Describe("Validate the create function", func() {
		It("Fails when first is not provided", func() {
			_, err := useCase.Create("", "a@b.c")
			//Expect(member).To(Equal(nil))
			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(Equal("rpc error: code = InvalidArgument desc = first not provided"))
		})
		It("Fails when email is not provided", func() {
			_, err := useCase.Create("mario", "")
			//Expect(member).To(Equal(nil))
			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(Equal("rpc error: code = InvalidArgument desc = email not provided"))
		})
		It("Fails when email is invalid", func() {
			_, err := useCase.Create("mario", "a@b")
			//Expect(member).To(not(Equal(nil))
			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(Equal("rpc error: code = InvalidArgument desc = invalid email address"))
		})
		It("Works when first and email are provided and valid", func() {
			var first = "mario"
			var email = "mario@yahoo.com"
			member, _ := useCase.Create(first, email)
			Expect(member).To(Not(BeNil()))
			Expect(member.First()).To(Equal(first))
			Expect(member.Email()).To(Equal(email))
		})
	})
	Describe("Validate the ReadId function", func() {
		It("Works when the id is valid", func() {
			var first = "mario"
			var email = "mario@yahoo.com"
			member, _ := useCase.Create(first, email)
			var id = member.Id()
			memberRead, _ := useCase.ReadById(id)
			Expect(memberRead).To(Not(BeNil()))
			Expect(memberRead.First()).To(Equal(first))
			Expect(memberRead.Email()).To(Equal(email))
			Expect(memberRead.Id()).To(Equal(id))
		})
		It("Fails when the id is valid", func() {
			var first = "mario"
			var email = "mario@yahoo.com"
			useCase.Create(first, email)
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
			var first = "mario"
			var email = "mario@yahoo.com"
			member, _ := useCase.Create(first, email)
			var id = member.Id()
			memberRead, _ := useCase.ReadByEmail(email)
			Expect(memberRead).To(Not(BeNil()))
			Expect(memberRead.First()).To(Equal(first))
			Expect(memberRead.Email()).To(Equal(email))
			Expect(memberRead.Id()).To(Equal(id))
		})
		It("Fails when the email is invalid", func() {
			var first = "mario"
			var email = "mario@yahoo.com"
			useCase.Create(first, email)
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
			member, error := useCase.Create(first, email)
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
			useCase.Create("mario", "mario@yahoo.com")
			useCase.Create("maria", "maria@yahoo.com")
			// read all members
			members, err := useCase.Read()
			Expect(members).To(Not(BeNil()))
			Expect(err).To(BeNil())
			Expect(len(members)).To(Equal(2))
			member := members[0]
			Expect(member.First()).To(Equal("mario"))
			Expect(member.Email()).To(Equal("mario@yahoo.com"))
			member = members[1]
			Expect(member.First()).To(Equal("maria"))
			Expect(member.Email()).To(Equal("maria@yahoo.com"))
		})

	})
})
