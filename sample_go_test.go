package sample

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sample", func() {
	Describe("Calculating square root", func() {
		Context("of the number 4", func() {
			It("should result 1", func() {
				Setup()
				Expect(GetResult()).To(Equal(2.0))
			})
		})
	})
})
