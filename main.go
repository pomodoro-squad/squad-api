package main

import (
	"flag"
	"github.com/toma-san/squad-api/server"
	"github.com/toma-san/squad-api/utils/application"
)

func main() {
	configFile := flag.String("config", "env/config.json", "Default config")
	firebaseFile := flag.String("firebase", "env/toma-san-firebase.json", "Default config for firebase")
	flag.Parse()
	app := application.Application{}
	app.InitConfiguration(*configFile)
	app.InitFirebase(*firebaseFile)
	server.StartServer(app)
}
