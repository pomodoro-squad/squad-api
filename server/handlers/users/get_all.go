package users

import (
	"github.com/labstack/echo"
	"net/http"

	m "github.com/toma-san/squad-api/models"
	u "github.com/toma-san/squad-api/utils"
)

func (h ApiV1Handler) GetAllUsersHandler(c echo.Context) error {
	users:= m.UserList{}
	if err := h.Firebase.Value(&users.Users); err != nil {
		u.SendError(http.StatusInternalServerError, c, err, ErrInternalDatabase)
	}

	return c.JSON(http.StatusOK, users)
}
