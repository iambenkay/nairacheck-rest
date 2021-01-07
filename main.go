package main

import (
	"context"
	"fmt"
	"github.com/iambenkay/nairacheck/controllers"
	"github.com/iambenkay/nairacheck/services"
	"github.com/iambenkay/nairacheck/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net"
	"net/http"
)

func main() {
	go services.InitializeDatabaseConnection(MongoURI)
	defer utils.Contextualize(func(ctx context.Context) {
		if err := services.Bean.DatabaseClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	})

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
