package main

import (
	"os"

	"github.com/manyminds/soyfr/library/server"
)

func main() {
	app := server.GetApplication()
	app.Run(os.Args)
}
