package application

import (
	"flag"

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

func (a *Application) InitConfiguration() {
	configfile := flag.String("config", "config.json", "Config for connection to database")
	flag.Parse()
	a.Configuration = config.MustNewConfig(*configfile)
	fmt.Println(a.Configuration)
}

func (a *Application) InitFirebase() {
	d, err := ioutil.ReadFile("toma-san-firebase-adminsdk-ogt0j-2b156ba741.json")
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

	a.Firebase = firego.New("https://toma-san.firebaseio.com/", conf.Client(context.Background()))
}