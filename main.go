package main

import (
	"net/http"

	"github.com/manyminds/soyfer/library"
	"github.com/maxwellhealth/bongo"
	"github.com/univedo/api2go"
)

func main() {
	//TODO configure via environment variables
	config := bongo.Config{
		ConnectionString: "localhost",
		Database:         "soyfer",
	}

	connection, err := bongo.Connect(&config)
	if err != nil {
		panic(err)
	}

	//TODO hide connection in construct
	userSource := library.UserSource{Connection: connection}

	api := api2go.NewAPI("v1")
	api.AddResource(library.User{}, &userSource)
	http.ListenAndServe(":8080", api.Handler())

}
