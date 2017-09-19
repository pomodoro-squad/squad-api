package users

import (
	"github.com/toma-san/squad-api/config"
	"github.com/acoshift/go-firebase-admin"
)

type ApiV1Handler struct {
	Configuration *config.Configuration
	Firebase *firebase.App
}

var (
	ErrInternalDatabase       string = "Database internal error"
	ErrInvalidUserId          string = "Id should be numeric"
	ErrIncorrectData          string = "Incorrect data or format of data sent"
	StatusUnprocessableEntity int    = 422
)