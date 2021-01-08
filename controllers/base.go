package controllers

import (
	"github.com/iambenkay/nairacheck/services"
	"github.com/labstack/echo/v4"
)

func Load() {
	services.Bean.Server.GET(FraudResource, FetchFraudsters())
	services.Bean.Server.GET(FraudResourceItem, FetchSingleFraudster())
	services.Bean.Server.POST(FraudResource, AddFraudster())
	services.Bean.Server.GET(Version, func(c echo.Context) error {
		return c.String(200, "0.1.0-unstable")
	})
}

type response struct {
	Message string      `json:"message"`
	Error   bool        `json:"error"`
	Data    interface{} `json:"data"`
}
