package users

import (
	"github.com/tomo-san/squad-api/config"
	"github.com/zabawaba99/firego"
)

type ApiV1Handler struct {
	Configuration *config.Configuration
	Firebase *firego.Firebase
}

var (
	ErrInternalDatabase       string = "Database internal error"
	ErrInvalidUserId          string = "Id should be numeric"
	ErrIncorrectData          string = "Incorrect data or format of data sent"
	StatusUnprocessableEntity int    = 422
)