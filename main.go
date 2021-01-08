package main

import (
	"fmt"
	"github.com/iambenkay/nairacheck/controllers"
	"github.com/iambenkay/nairacheck/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net"
	"net/http"
)

func main() {
	go services.InitializeDatabaseConnection(MongoURI)
	defer services.DestroyDatabaseConnection()

	e := echo.New()

	services.Bean.Server = e

	e.Use(middleware.Logger())
	listener, err := net.Listen("tcp4", fmt.Sprintf(":%s", env("PORT")))
	if err != nil {
		panic(err)
	}

	e.Listener = listener
	s := new(http.Server)

	controllers.Load()

	err = e.StartServer(s)
	if err != nil {
		panic(err)
	}
}
