package users

import (
	"github.com/labstack/echo"
	"net/http"

	m "github.com/toma-san/squad-api/models"
	u "github.com/toma-san/squad-api/utils"
	"fmt"
)

func (h ApiV1Handler) GetAllUsersHandler(c echo.Context) error {
	users:= []m.User{}

	if err := h.Firebase.Child("users").Value(&users); err != nil {
		u.SendError(http.StatusInternalServerError, c, err, ErrInternalDatabase)
		fmt.Println(err)
	}
	users = users[1:]


	return c.JSON(http.StatusOK, users)
}
