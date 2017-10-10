package application

import (
	"github.com/toma-san/squad-api/config"
	"github.com/zabawaba99/firego"
	"io/ioutil"
	"golang.org/x/oauth2/google"
	"golang.org/x/net/context"
	"fmt"
)

type App interface {
	InitConfiguration()
	InitFirebase()
}

type Application struct {
	Configuration *config.Configuration
	Firebase *firego.Firebase
}

func (a *Application) InitConfiguration(configfile string) {
	a.Configuration = config.MustNewConfig(configfile)
}

func (a *Application) InitFirebase(configfile string) {
	d, err := ioutil.ReadFile(configfile)
	if err != nil {
		fmt.Errorf("Error appeared: %s", err)
		//return nil, err
	}
	conf, err := google.JWTConfigFromJSON(d, "https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/firebase.database")
	if err != nil {
		fmt.Errorf("Error appeared: %s", err)
		//return nil, err
	}

	a.Firebase = firego.New(a.Configuration.FirebaseConfig.URl, conf.Client(context.Background()))
}