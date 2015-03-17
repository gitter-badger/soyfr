package main

import (
	"net/http"

	"github.com/manyminds/soyfr/library"
	"github.com/maxwellhealth/bongo"
)

func main() {
	//TODO configure via environment variables
	config := bongo.Config{
		ConnectionString: "localhost",
		Database:         "soyfer",
	}

	http.ListenAndServe(":8080", library.BootstrapAPI(&config))
}
