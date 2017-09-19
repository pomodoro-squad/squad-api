package application

import (
	"flag"

	"github.com/toma-san/squad-api/config"
	f "github.com/toma-san/squad-api/firebase"
	"github.com/acoshift/go-firebase-admin"
)

type App interface {
	InitConfiguration()
	InitFirebase()
}

type Application struct {
	Configuration *config.Configuration
	Firebase *firebase.App
}

func (a *Application) InitConfiguration() {
	configfile := flag.String("config", "config.json", "Config for connection to database")
	flag.Parse()
	a.Configuration = config.MustNewConfig(*configfile)
}

func (a *Application) InitFirebase() {
	a.Firebase = f.InitFire(a.Configuration)
	a.Firebase.Auth()
}