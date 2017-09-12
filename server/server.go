package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	a "github.com/toma-san/squad-api/utils/application"
)

func StartServer(application a.Application){
	server := echo.New()

	server.Use(middleware.Recover())

	server.GET("/status", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Status ok")
	})
	server.Logger.Fatal(server.Start(fmt.Sprintf(":%d", application.Configuration.ListenAddress)))
}