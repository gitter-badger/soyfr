package library_test

import (
	. "github.com/manyminds/soyfr/library"
	"github.com/maxwellhealth/bongo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/univedo/api2go"
)

var _ = Describe("User", func() {
	var connection *bongo.Connection
	var userSource UserSource
	var request api2go.Request

	BeforeEach(func() {
		var err error
		config := bongo.Config{
			ConnectionString: "localhost",
			Database:         "soyfer_test",
		}

		connection, err = bongo.Connect(&config)
		Expect(err).ToNot(HaveOccurred())

		userSource = UserSource{Connection: connection}
	})

	Context("basic user crud api methods", func() {
		It("Should create a new user", func() {
			By("storing it")
			user := User{Username: "Unittest"}
			id, err := userSource.Create(user)
			Expect(err).ToNot(HaveOccurred())
			Expect(id).ToNot(Equal(""))
			By("finding it again")
			after, err := userSource.FindOne(id, request)
			Expect(err).ToNot(HaveOccurred())
			castedUser, ok := after.(User)
			Expect(ok).To(Equal(true))
			Expect(id).To(Equal(castedUser.GetId().Hex()))
		})

		It("Should find zero users", func() {
			resultSet, err := userSource.FindAll(request)
			Expect(err).ToNot(HaveOccurred())

			data, ok := resultSet.([]User)
			Expect(ok).To(Equal(true))
			Expect(data).To(HaveLen(0))
		})

		It("Should find all added users", func() {
			usersToAdd := []string{"userA", "userB", "userC"}
			for _, username := range usersToAdd {
				user := User{Username: username}
				_, err := userSource.Create(user)
				Expect(err).ToNot(HaveOccurred())
			}

			resultSet, err := userSource.FindAll(request)
			Expect(err).ToNot(HaveOccurred())

			data, ok := resultSet.([]User)
			Expect(ok).To(Equal(true))
			Expect(data).To(HaveLen(3))
		})
	})

	AfterEach(func() {
		connection.Session.DB("soyfer_test").DropDatabase()
	})
})
