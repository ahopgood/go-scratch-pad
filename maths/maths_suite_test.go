package maths_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMaths(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Maths Suite")
}
