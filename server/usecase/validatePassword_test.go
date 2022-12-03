package usecase_test

import (
	uc "github.com/RodrigoMattosoSilveira/go-swiss-pairing/server/usecase"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validate the password validator", func() {
	It("works for a short password", func() {
		Expect(uc.ValidatePassword("oPT14J")).To(BeFalse())
	})
	It("works for a long password", func() {
		Expect(uc.ValidatePassword("oPT14J30I#y4oPT14J30I#y4oPT14J30I#y4")).To(BeFalse())
	})
	It("works for a password without any digits", func() {
		Expect(uc.ValidatePassword("oPTIgJEOI#yF")).To(BeFalse())
	})
	It("works for a password without any lower case characters", func() {
		Expect(uc.ValidatePassword("OPT14J30I#Y4#YF")).To(BeFalse())
	})
	It("works for a password without any upper case characters", func() {
		Expect(uc.ValidatePassword("opt14j30i#y4")).To(BeFalse())
	})
	It("works with at least one space", func() {
		Expect(uc.ValidatePassword("o T14J30I#y4")).To(BeFalse())
	})
	It("works for a valid password", func() {
		Expect(uc.ValidatePassword("oPT14J30I#y4")).To(BeTrue())
	})
})
