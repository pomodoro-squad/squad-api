package users

import (
	"github.com/labstack/echo"
	"net/http"

	//m "github.com/toma-san/squad-api/models"
	//u "github.com/toma-san/squad-api/utils"
	//"github.com/tomo-san/squad-api/server/handlers/users"
)

func (h ApiV1Handler) GetAllUsersHandler(c echo.Context) error {
	//users:= m.UserList{}
	users := h.Firebase.Auth().ListUsers(100)
	//if err != nil {
	//	u.SendError(http.StatusInternalServerError, c, err, ErrInternalDatabase)
	//}

	return c.JSON(http.StatusOK, users)
}
