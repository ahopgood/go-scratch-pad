package maths_test

import (
	"com/alexander/scratch/maths"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var fibonacci func() int

var _ = Describe("Fibonacci", func() {

	BeforeEach(func() {
		fibonacci = maths.Fibonacci()
	})

	
	Context("Testing", func(){
		It("Should pass", func() {
			Expect("Foo").To(Equal("Foo"))
		})
	})
	Context("when on one iteration", func() {
		It("returns 0", func() {
			Expect(fibonacci()).To(Equal(0))
		})
	})

	Context("when on second iteration", func() {
		It("returns 1", func() {
			fibonacci()
			Expect(fibonacci()).To(Equal(1))
		})
	})

	Context("when on third iteration", func() {
		It("returns 1", func() {
			fibonacci()
			fibonacci()
			Expect(fibonacci()).To(Equal(1))
		})
	})

	Context("when on fourth iterations", func() {
		It("returns 2", func() {
			fibonacci()
			fibonacci()
			fibonacci()
			Expect(fibonacci()).To(Equal(2))
		})
	})

	Context("when on fifth iterations", func() {
		It("returns 3", func() {
			fibonacci()
			fibonacci()
			fibonacci()
			fibonacci()
			Expect(fibonacci()).To(Equal(3))
		})
	})
})
