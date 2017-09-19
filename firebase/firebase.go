package firebase

import (
	"golang.org/x/net/context"

	"github.com/toma-san/squad-api/config"
	"google.golang.org/api/option"
	"github.com/acoshift/go-firebase-admin"
)

func InitFirebase (c config.Configuration){
	// Init App with service_account
	option.WithCredentialsFile(c.FirebaseConfig.CredentialFile)

	firApp, err := firebase.InitializeApp(context.Background(), firebase.AppOptions{
		ProjectID:      "toma-san",
	}, option.WithCredentialsFile(c.FirebaseConfig.CredentialFile))

	if err != nil {
		panic(err)
	}

	//fd := firApp.Database()
	_ := firApp.Database()


}