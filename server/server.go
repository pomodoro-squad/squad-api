package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	a "github.com/toma-san/squad-api/utils/application"
	"github.com/toma-san/squad-api/server/handlers/users"
	mw "github.com/toma-san/squad-api/middleware"
)

func StartServer(application a.Application){
	server := echo.New()

	userConfig := users.ApiV1Handler{
		Configuration: application.Configuration,
		Firebase: application.Firebase,
	}

	server.Use(mw.Logger())
	server.Use(middleware.Recover())


	server.GET("/status", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Status ok")
	})
	v1 := server.Group("/v1")

	user := v1.Group("/users")

	user.GET("", userConfig.GetAllUsersHandler)
	user.GET("/:id", userConfig.GetUserHandler)

	server.Logger.Fatal(server.Start(fmt.Sprintf(":%d", application.Configuration.ListenAddress)))
}