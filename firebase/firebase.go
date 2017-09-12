package firebase

import (
	"github.com/toma-san/squad-api/config"
	"github.com/zabawaba99/firego"
)

func NewFirebase(c config.FirebaseConfig) *firego.Firebase{
	f := firego.New("https://toma-san.firebaseio.com", nil)
	return f
}