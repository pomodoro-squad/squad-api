package utils

import (
	"github.com/tomo-san/squad-api/config"
	"fmt"
)

func GetTocken(c config.Configuration)(string, error) {
	client, err := app.Auth()
	if err != nil {
		return "", fmt.Errorf("error getting Auth client: %v", err)
	}

	claims := map[string]interface{}{
		"premiumAccount": true,
	}

	token, err := client.CustomTokenWithClaims("some-uid", claims)
	if err != nil {
		return "", fmt.Errorf("error minting custom token: %v", err)
	}

	fmt.Printf("Got custom token: %v\n", token)
}