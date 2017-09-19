package users


import (
	"net/http"

	"github.com/labstack/echo"
	//m "github.com/toma-san/squad-api/models"
	u "github.com/toma-san/squad-api/utils"
	"context"
)

func (h ApiV1Handler) GetUserHandler(c echo.Context) error {
	//user := m.User{}
	var id []string
	id[0] = c.Param("id")
	user, err := h.Firebase.Auth().GetUsers(context.Background(), id)


	if err != nil {
		u.SendError(http.StatusInternalServerError, c, err, ErrInternalDatabase)
	}
	return c.JSON(http.StatusOK, user)
}
