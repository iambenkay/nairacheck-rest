package services

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceGroup struct {
	Server *echo.Echo
	DatabaseClient *mongo.Client
}

var Bean ServiceGroup