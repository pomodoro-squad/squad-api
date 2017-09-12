package users


import (
	"net/http"

	"github.com/labstack/echo"
	m "github.com/toma-san/squad-api/models"
	u "github.com/toma-san/squad-api/utils"
)

func (h ApiV1Handler) GetUserHandler(c echo.Context) error {
	user := m.User{}

	if err := h.Firebase.EqualTo(c.Param("id")).Value(&user); err != nil {
		u.SendError(http.StatusInternalServerError, c, err, ErrInternalDatabase)
	}
	return c.JSON(http.StatusOK, user)
}
