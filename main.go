package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/lucas-clemente/go-http-logger"
	"github.com/manyminds/soyfr/library"
	"github.com/maxwellhealth/bongo"
)

//wrapFileHandler adds a wildcard to index.html if there are
//static routes defined such as /fish or /some-seo-route
//files with extension will still be delivered normally.
func wrapFileHandler(distPath string, handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, ".") {
			handler.ServeHTTP(w, r)
			return
		}

		filename := distPath + r.URL.Path

		if _, err := os.Stat(filename); os.IsNotExist(err) {
			index, err := ioutil.ReadFile(distPath + "/index.html")
			if err != nil {
				log.Println("Could not find index.html")
				w.WriteHeader(404)
				return
			}

			w.Write(index)
			return
		}

		handler.ServeHTTP(w, r)
	}
}

//wrapAPIHandler is a hack to let api2go be used within
//the normal http mux functionality of go
func wrapAPIHandler(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.Replace(r.RequestURI, "/api", "", 1)
		handler.ServeHTTP(w, r)
	}
}

func main() {
	//TODO configure via environment variables
	config := bongo.Config{
		ConnectionString: "localhost",
		Database:         "soyfr",
	}

	distPath := "./app/"

	mux := http.NewServeMux()
	fileHandler := http.FileServer(http.Dir(distPath))
	mux.Handle("/api/", wrapAPIHandler(library.BootstrapAPI(&config)))
	mux.Handle("/", wrapFileHandler(distPath, fileHandler))

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", logger.Logger(mux))
}
