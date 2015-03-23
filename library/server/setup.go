package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/lucas-clemente/go-http-logger"
	"github.com/manyminds/soyfr/library/db"
	"github.com/maxwellhealth/bongo"
)

const (
	//EnvServerPort constant contains the name for the env variable to define the port
	//where the server listens.
	EnvServerPort = "SOYFR_SERVER_PORT"
	//EnvConnectionURI is the mgo uri to the database or replica set
	EnvConnectionURI = "SOYFR_CONNECTION_URI"
	//EnvDatabase constant is the mongo database name for this application.
	EnvDatabase = "SOYFR_DATABASE"
	//EnvResourceFiles is the relative or absolute path to the folder where
	//the static frontend files are
	EnvResourceFiles = "SOYFR_RESOURCE_FILES"
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

//GetApplication will build an cli application
//with basic flags for env, or bash params
func GetApplication() *cli.App {
	app := cli.NewApp()
	app.Name = "Soyfr"
	app.Usage = "A crowd based party drinking game"
	app.Author = "Manyminds"
	app.Email = "info@manyminds.de"
	app.Version = "Development"

	serverPortFlag := cli.IntFlag{
		Name:   "port",
		Value:  8800,
		Usage:  "9000",
		EnvVar: EnvServerPort,
	}

	databaseString := cli.StringFlag{
		Name:   "database",
		Value:  "soyfr_development",
		Usage:  "some name for a database",
		EnvVar: EnvDatabase,
	}

	distPathString := cli.StringFlag{
		Name:   "resourceDirectory",
		Value:  "./app",
		Usage:  "path to the resource files",
		EnvVar: EnvResourceFiles,
	}

	app.Flags = []cli.Flag{serverPortFlag, databaseString, distPathString}
	app.Action = func(c *cli.Context) {
		connectionString := db.GetConnectionString()
		database := c.String("database")
		distPath := c.String("resourceDirectory")
		serverPort := c.Int("port")

		log.Printf("Mongo connection on %s\n", connectionString)

		startApplication(connectionString, database, distPath, serverPort)
	}

	return app
}

func startApplication(connectionString, database, distPath string, serverPort int) {
	config := bongo.Config{
		ConnectionString: connectionString,
		Database:         database,
	}

	mux := http.NewServeMux()
	fileHandler := http.FileServer(http.Dir(distPath))
	mux.Handle("/api/", wrapAPIHandler(db.BootstrapAPI(&config)))
	mux.Handle("/", wrapFileHandler(distPath, fileHandler))

	log.Printf("Server started on port :%d\n", serverPort)
	http.ListenAndServe(fmt.Sprintf(":%d", serverPort), logger.Logger(mux))
}
