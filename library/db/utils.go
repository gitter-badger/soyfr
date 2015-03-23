package db

import (
	"net/http"
	"os"
	"strings"

	"github.com/maxwellhealth/bongo"
	"github.com/univedo/api2go"
)

const (
	//DockerConnectionString this string connects to a docker link TODO normal configuration
	DockerConnectionString = "MONGODB_PORT_27017_TCP"
	//DirectConnectionString is used when connecting without docker
	DirectConnectionString = "MONGO_CONNECTION_STRING"
	//FallbackConnectionString is the one we use if everthing else fails
	FallbackConnectionString = "localhost:27017"
)

//GetConnectionString returns the configured connection string
func GetConnectionString() string {
	dockerEnv := os.Getenv(DockerConnectionString)
	if dockerEnv != "" {
		return strings.Replace(dockerEnv, "tcp://", "", 1)
	}

	normalEnv := os.Getenv(DirectConnectionString)
	if normalEnv != "" {
		return normalEnv
	}

	return FallbackConnectionString
}

//BootstrapAPI configures the api and returns the corresponding handler
func BootstrapAPI(config *bongo.Config) http.Handler {
	userSource, err := CreateUserSource(config)

	if err != nil {
		panic(err)
	}

	api := api2go.NewAPI("v1")
	api.AddResource(User{}, userSource)

	return api.Handler()
}
