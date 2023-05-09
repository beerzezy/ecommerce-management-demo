package authentication

import (
	"net/http"

	"github.com/beerzezy/ecommerce-management-demo/internal/baseresponse"
	"github.com/beerzezy/ecommerce-management-demo/internal/errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Handler interface {
	Login(c echo.Context) error
	Registration(c echo.Context) error
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return handler{service}
}

func (h handler) Login(c echo.Context) error {

	var request RequestLogin
	if err := c.Bind(&request); err != nil {
		log.Errorf("Login bind baseRequest: %v", err)
		return err
	}

	if err := c.Validate(request); err != nil {
		log.Errorf("Login Validate Request: %v", err)
		c.JSON(http.StatusOK, errors.NewExceptionError(err.Error()))
		return err
	}

	result, err := h.service.Login(request)
	if err != nil {
		log.Errorf("Login Error: %v", err)
		c.JSON(http.StatusOK, errors.NewExceptionError(err.Error()))
		return err
	}

	return c.JSON(http.StatusOK, baseresponse.BaseResponse{
		Message: "success",
		Data:    result,
	})
}

func (h handler) Registration(c echo.Context) error {

	var request RequestRegistration
	if err := c.Bind(&request); err != nil {
		log.Errorf("Registration bind baseRequest: %v", err)
		return err
	}

	if err := c.Validate(request); err != nil {
		log.Errorf("Registration Validate Request: %v", err)
		return errors.NewInvalidRequireFieldError()
	}

	err := h.service.Registration(request)
	if err != nil {
		log.Errorf("Registration Error: %v", err)
		return c.JSON(http.StatusOK, errors.NewExceptionError(err.Error()))
	}

	return c.JSON(http.StatusOK, baseresponse.BaseResponse{
		Message: "success",
	})
}
