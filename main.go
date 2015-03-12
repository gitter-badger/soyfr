package main

import (
	"net/http"

	"github.com/manyminds/soyfr/library"
	"github.com/maxwellhealth/bongo"
	"github.com/univedo/api2go"
)

func main() {
	//TODO configure via environment variables
	config := bongo.Config{
		ConnectionString: "localhost",
		Database:         "soyfer",
	}

	usersource, err := library.CreateUserSource(&config)

	if err != nil {
		panic(err)
	}

	api := api2go.NewAPI("v1")
	api.AddResource(library.User{}, userSource)
	http.ListenAndServe(":8080", api.Handler())
}
