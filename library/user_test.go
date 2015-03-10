package library_test

import (
	"github.com/manyminds/soyfer/library"
	"github.com/maxwellhealth/bongo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/univedo/api2go"
)

var _ = Describe("User", func() {
	var connection *bongo.Connection
	var userSource library.UserSource

	BeforeEach(func() {
		var err error
		config := bongo.Config{
			ConnectionString: "localhost",
			Database:         "soyfer_test",
		}

		connection, err = bongo.Connect(&config)
		Expect(err).ToNot(HaveOccurred())

		userSource = library.UserSource{Connection: connection}
	})

	Context("basic user crud api methods", func() {
		It("Should create a new user", func() {
			By("storing it")
			user := library.User{Username: "Unittest"}
			id, err := userSource.Create(user)
			Expect(err).ToNot(HaveOccurred())
			Expect(id).ToNot(Equal(""))
			By("finding it again")
			after, err := userSource.FindOne(id, api2go.Request{})
			Expect(err).ToNot(HaveOccurred())
			castedUser, ok := after.(library.User)
			Expect(ok).To(Equal(true))
			Expect(id).To(Equal(castedUser.GetId().Hex()))
		})
	})

	AfterEach(func() {
		connection.Session.DB("soyfer_test").DropDatabase()
	})
})
