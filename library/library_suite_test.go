package library_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestLibrary(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Library Suite")
}
