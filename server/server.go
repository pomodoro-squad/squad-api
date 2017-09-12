package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	a "github.com/toma-san/squad-api/utils/application"
	"github.com/toma-san/squad-api/server/handlers/users"
)

func StartServer(application a.Application){
	server := echo.New()

	server.Use(middleware.Recover())

	userConfig := users.ApiV1Handler{}

	server.GET("/status", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Status ok")
	})
	v1 := server.Group("/v1")

	user := v1.Group("/users")

	user.GET("", userConfig.GetUserHandler)
	user.GET("/:id", userConfig.GetAllUsersHandler)

	server.Logger.Fatal(server.Start(fmt.Sprintf(":%d", application.Configuration.ListenAddress)))
}