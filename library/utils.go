package library

import (
	"net/http"

	"github.com/maxwellhealth/bongo"
	"github.com/univedo/api2go"
)

//bootstrapAPI configures the api and returns the corresponding handler
func bootstrapAPI(config *bongo.Config) http.Handler {
	userSource, err := CreateUserSource(config)

	if err != nil {
		panic(err)
	}

	api := api2go.NewAPI("v1")
	api.AddResource(User{}, userSource)

	return api.Handler()
}