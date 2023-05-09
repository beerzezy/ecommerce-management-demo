package usermanagement

import (
	"net/http"

	"github.com/beerzezy/ecommerce-management-demo/internal/baseresponse"
	"github.com/beerzezy/ecommerce-management-demo/internal/errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Handler interface {
	Registration(c echo.Context) error
	GetAccountInfo(c echo.Context) error
	GetAccountInfoAll(c echo.Context) error
	GetAccountInfoByID(c echo.Context) error
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return handler{service}
}

func (h handler) GetAccountInfoByID(c echo.Context) error {
	return nil
}

func (h handler) GetAccountInfoAll(c echo.Context) error {
	return nil
}

func (h handler) GetAccountInfo(c echo.Context) error {

	var request RequestGetAccountInfo
	if err := c.Bind(&request); err != nil {
		log.Errorf("GetAccountInfo bind baseRequest: %v", err)
		return err
	}

	if err := c.Validate(request); err != nil {
		log.Errorf("GetAccountInfo Validate Request: %v", err)
		c.JSON(http.StatusOK, errors.NewExceptionError(err.Error()))
		return err
	}

	result, err := h.service.GetAccountInfo(request.Email)
	if err != nil {
		log.Errorf("GetAccountInfo Error: %v", err)
		c.JSON(http.StatusOK, errors.NewExceptionError(err.Error()))
		return err
	}

	return c.JSON(http.StatusOK, baseresponse.BaseResponse{
		Message: "success",
		Data:    result,
	})
}

func (h handler) Registration(c echo.Context) error {

	var request RequestCreateAccount
	if err := c.Bind(&request); err != nil {
		log.Errorf("Registration bind baseRequest: %v", err)
		return err
	}

	if err := c.Validate(request); err != nil {
		log.Errorf("Registration Validate Request: %v", err)
		c.JSON(http.StatusOK, errors.NewExceptionError(err.Error()))
		return err
	}

	err := h.service.CreateAccount(request)
	if err != nil {
		log.Errorf("Registration Error: %v", err)
		return c.JSON(http.StatusOK, errors.NewExceptionError(err.Error()))
	}

	return c.JSON(http.StatusOK, baseresponse.BaseResponse{
		Message: "success",
	})
}
