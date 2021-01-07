package controllers

import (
	"github.com/iambenkay/nairacheck/models"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func FetchFraudsters() echo.HandlerFunc {
	return func(c echo.Context) error {
		params := new(models.PageSortParams)
		if err := c.Bind(params); err != nil {
			c.Error(
				&echo.HTTPError{
					Message:  "bind error occurred",
					Code:     400,
					Internal: err,
				})
			return nil
		}
		params.Paged = true
		fraudsters, err := models.FindFraudsters(bson.M{}, params)

		if err != nil {
			c.Error(
				&echo.HTTPError{
					Message:  err.Error(),
					Code:     400,
					Internal: err,
				})
			return nil
		}

		return c.JSON(200, response{
			Message: "Fraudsters fetched successfully",
			Error:   false,
			Data:    fraudsters,
		})
	}
}
