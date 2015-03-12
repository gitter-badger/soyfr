package library_test

import (
	"github.com/maxwellhealth/bongo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func getDatabaseConfiguration() *bongo.Config {
	//TODO configure via environment variables
	return &bongo.Config{
		ConnectionString: "localhost",
		Database:         "soyfer_test",
	}
}

func TestLibrary(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Library Suite")
}
