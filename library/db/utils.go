package db

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/googollee/go-socket.io"
	"github.com/manyminds/api2go"
	"github.com/maxwellhealth/bongo"
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

//BootstrapAPI
func BootstrapAPI(config *bongo.Config) http.Handler {
	api := api2go.NewAPI("v1")

	return api.Handler()
}

//BootstrapWebsocket configures the api and returns the corresponding handler
func BootstrapWebsocket(config *bongo.Config) http.Handler {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")
		so.Join("chat")
		so.On("disconnection", func() {
			log.Println("on disconnect")
			so.BroadcastTo("chat", "chat message", "fick dich")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	return server
}
