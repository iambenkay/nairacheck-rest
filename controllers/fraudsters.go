package controllers

import (
	"github.com/iambenkay/nairacheck/models"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FetchSingleFraudster() echo.HandlerFunc {
	return func(c echo.Context) error {
		filters := new(models.Fraudster)
		bindErr := echo.PathParamsBinder(c).CustomFunc("id", func(values []string) []error {
			var err error
			filters.ID, err = primitive.ObjectIDFromHex(values[0])
			return []error{err}
		}).BindError()

		if bindErr != nil {
			return &echo.HTTPError{
				Message:  bindErr.Error(),
				Code:     400,
				Internal: bindErr,
			}
		}

		fraudster, err := models.FindOneFraudster(*filters)
		if err != nil {
			return &echo.HTTPError{
				Message:  err.Error(),
				Code:     400,
				Internal: err,
			}
		}
		if fraudster == nil {
			return c.JSON(404, response{
				Message: "Fraudster was not found",
				Error:   true,
				Data:    nil,
			})
		}
		return c.JSON(200, response{
			Message: "Fraudster fetched successfully",
			Error:   false,
			Data:    fraudster,
		})
	}
}

func FetchFraudsters() echo.HandlerFunc {
	return func(c echo.Context) error {
		params := new(models.PageSortParams)
		binder := echo.DefaultBinder{}
		if err := binder.BindQueryParams(c, params); err != nil {
			return err
		}
		params.Paged = true

		filters := models.Fraudster{}
		if err := binder.BindQueryParams(c, &filters); err != nil {
			return err
		}
		fraudsters, err := models.FindFraudsters(filters, params)

		if err != nil {
			return &echo.HTTPError{
				Message:  err.Error(),
				Code:     400,
				Internal: err,
			}
		}

		return c.JSON(200, response{
			Message: "Fraudsters fetched successfully",
			Error:   false,
			Data:    fraudsters,
		})
	}
}

func AddFraudster() echo.HandlerFunc {
	return func(c echo.Context) error {
		binder := echo.DefaultBinder{}

		fraudster := models.Fraudster{}
		if err := binder.BindBody(c, &fraudster); err != nil {
			return err
		}
		*fraudster.Verified = false
		fraudster, err := models.AddFraudster(fraudster)

		if err != nil {
			return &echo.HTTPError{
				Message:  err.Error(),
				Code:     400,
				Internal: err,
			}
		}

		return c.JSON(200, response{
			Message: "Fraudsters added successfully",
			Error:   false,
			Data:    fraudster,
		})
	}
}
