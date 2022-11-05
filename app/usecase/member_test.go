package usecase_test

import (
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/domain/service"
	repo "github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/interface/persistence/memory"
	uc "github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/usecase"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
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
			Expect(err).To(Not(Equal(nil)))
			Expect(err.Error()).To(Equal("rpc error: code = InvalidArgument desc = first not provided"))
		})
		It("Fails when email is not provided", func() {
			_, err := useCase.Create("mario", "")
			//Expect(member).To(Equal(nil))
			Expect(err).To(Not(Equal(nil)))
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
})
