package usecase_test

import (
	uc "github.com/RodrigoMattosoSilveira/go-swiss-pairing/server/usecase"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validate the USA phone validator", func() {
	It("works for a phone missing the first 3 digits", func() {
		Expect(uc.ValidateUsaCellPhoneNumber("290 5243")).To(BeFalse())
	})
	It("works for a phone missing the second 3 digits", func() {
		Expect(uc.ValidateUsaCellPhoneNumber("208 5243")).To(BeFalse())
	})
	It("works for a phone missing the last 4 digits", func() {
		Expect(uc.ValidateUsaCellPhoneNumber("208 290")).To(BeFalse())
	})
	It("works for a phone with non digits", func() {
		Expect(uc.ValidateUsaCellPhoneNumber("208 290 524e")).To(BeFalse())
	})
	It("works for a bare valid phone number", func() {
		Expect(uc.ValidateUsaCellPhoneNumber("208 290 5243")).To(BeTrue())
	})
	It("works for a bare valid phone number without spaces", func() {
		Expect(uc.ValidateUsaCellPhoneNumber("2082905243")).To(BeTrue())
	})
	It("works for valid phone number with parenthesis around the first 3 digits", func() {
		Expect(uc.ValidateUsaCellPhoneNumber("(208) 290 5243")).To(BeTrue())
	})
	It("works for valid phone number with parenthesis around the first 3 digits and a dash between the 2nd and third groups", func() {
		Expect(uc.ValidateUsaCellPhoneNumber("(208) 290-5243")).To(BeTrue())
	})
	It("works for valid phone number without a space between the first and second group", func() {
		Expect(uc.ValidateUsaCellPhoneNumber("(208)290-5243")).To(BeTrue())
	})
})
