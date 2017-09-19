package firebase

import (
	"google.golang.org/api/option"
	"github.com/acoshift/go-firebase-admin"
	"github.com/toma-san/squad-api/config"
	"context"
)

func InitFire(c *config.Configuration)(*firebase.App){
	app, err := firebase.InitializeApp(context.Background(),
		firebase.AppOptions{
			ProjectID: "Toma-san",
		},
			option.WithCredentialsFile(c.FirebaseConfig.CredentialFile))
	if err != nil {
		panic(err)
	}

	return app
}