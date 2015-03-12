package library

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/maxwellhealth/bongo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/univedo/api2go"
	"gopkg.in/mgo.v2/bson"
)

var _ = Describe("User", func() {
	var userSource *UserSource
	var request api2go.Request

	getBodyFromURL := func(URL string) string {
		resp, err := http.Get(URL)
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))
		body, err := ioutil.ReadAll(resp.Body)
		Expect(err).ToNot(HaveOccurred())
		stringBody := string(body)
		return stringBody
	}

	BeforeEach(func() {
		rand.Seed(time.Now().UnixNano())
		var err error
		userSource, err = CreateUserSource(getDatabaseConfiguration())
		Expect(err).ToNot(HaveOccurred())
	})

	Context("test create via api", func() {
		var server *httptest.Server
		BeforeEach(func() {
			server = httptest.NewServer((bootstrapAPI(getDatabaseConfiguration())))
		})

		It("Should be able to list users", func() {
			Expect(getBodyFromURL(server.URL + "/v1/users")).To(MatchJSON("{\"data\":[]}"))
		})
	})

	Context("basic user crud model methods", func() {
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

		It("Should create a new user and update him", func() {
			By("storing it")
			user := User{Username: "Unittest"}
			id, err := userSource.Create(user)
			Expect(err).ToNot(HaveOccurred())
			Expect(id).ToNot(Equal(""))
			user.ID = bson.ObjectIdHex(id)

			By("renaming him")
			user.Username = "New Unittest"
			err = userSource.Update(user)
			Expect(err).ToNot(HaveOccurred())

			By("retrieving him from the database")
			after, err := userSource.FindOne(id, request)
			Expect(err).ToNot(HaveOccurred())
			castedUser, ok := after.(User)
			Expect(ok).To(Equal(true))
			Expect(id).To(Equal(castedUser.GetId().Hex()))
			Expect(castedUser.Username).To(Equal("New Unittest"))
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

		It("Should find some added users", func() {
			maxUsers := 100
			var idsToFind []string
			var i int

			for i < maxUsers {
				i++
				user := User{Username: fmt.Sprintf("user_%d", i)}
				idString, err := userSource.Create(user)
				Expect(err).ToNot(HaveOccurred())

				if rand.Int()%2 == 0 {
					idsToFind = append(idsToFind, idString)
				}
			}

			By(fmt.Sprintf("Finding %d users", len(idsToFind)))

			resultSet, err := userSource.FindMultiple(idsToFind, request)
			Expect(err).ToNot(HaveOccurred())

			data, ok := resultSet.([]User)
			Expect(ok).To(Equal(true))
			Expect(data).To(HaveLen(len(idsToFind)))
		})
	})

	AfterEach(func() {
		if con, err := bongo.Connect(getDatabaseConfiguration()); err == nil {
			con.Session.DB("soyfer_test").DropDatabase()
		}
	})
})
