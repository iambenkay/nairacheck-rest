package controllers

import (
	"github.com/iambenkay/nairacheck/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Load() {
	services.Bean.Server.GET(FraudResource, FetchFraudsters())
	services.Bean.Server.GET(FraudResourceItem, FetchSingleFraudster())
	services.Bean.Server.POST(FraudResource, AddFraudster(), BasicAuthMW)
	services.Bean.Server.GET(Version, func(c echo.Context) error {
		return c.String(200, "0.1.0-unstable")
	})
}

type response struct {
	Message string      `json:"message"`
	Error   bool        `json:"error"`
	Data    interface{} `json:"data"`
}

var BasicAuthMW = middleware.BasicAuth(func(username string, password string, c echo.Context) (bool, error) {
	if username == "core-service" && password == "root" {
		return true, nil
	}
	return false, c.JSON(401, response{
		Data:    nil,
		Message: "You are not authorized to access this route",
		Error:   true,
	})
})
