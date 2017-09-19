package application

import (
	"flag"

	"github.com/toma-san/squad-api/config"
	//"github.com/toma-san/squad-api/firebase"
	"github.com/zabawaba99/firego"
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
	a.Firebase = firego.New(a.Configuration.FirebaseConfig.URl, nil)
	a.Firebase.Unauth()
}