package application

import (
	"flag"

	"github.com/toma-san/squad-api/config"
	"github.com/zabawaba99/firego"
	"github.com/toma-san/squad-api/firebase"
)

type App interface {
	InitConfiguration()
	InitFirebase()
}

type Application struct {
	Configuration *config.Configuration
	Firebase *firego.Firebase
}

func (a *Application) InitConfiguration() {
	configfile := flag.String("config", "config.json", "Config for connection to database")
	flag.Parse()
	a.Configuration = config.MustNewConfig(*configfile)
}

func (a *Application) InitFirebase() {
	a.Firebase = firebase.NewFirebase(a.Configuration.FirebaseConfig)
	a.Firebase.Auth()
}