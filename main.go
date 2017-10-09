package main

import (
	"github.com/toma-san/squad-api/server"
	"github.com/toma-san/squad-api/utils/application"
)

func main() {
	app := application.Application{}
	app.InitConfiguration()
	app.InitFirebase()
	server.StartServer(app)
}