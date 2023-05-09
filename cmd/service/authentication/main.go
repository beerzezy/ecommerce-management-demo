package main

import (
	"fmt"
	"net/http"

	"github.com/beerzezy/ecommerce-management-demo/cmd/common"
	"github.com/beerzezy/ecommerce-management-demo/cmd/initialization"
	"github.com/beerzezy/ecommerce-management-demo/internal/config"
	"github.com/beerzezy/ecommerce-management-demo/internal/core/authentication"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func main() {
	config := initialization.InitConfig()
	db := initialization.InitDB(config)
	defer db.Close()

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	registerRoutes(e, db, config)

	go common.StartServer(e, config)
	common.WaitForGracefulShutdown(e)
}

func registerRoutes(e *echo.Echo, db *gorm.DB, cfg config.Config) {

	authentication := authentication.NewHandler(authentication.NewService(authentication.NewRepository(db)))

	e.GET("/health", health)

	serverContext := cfg.Server.Context
	loginPath := fmt.Sprintf("%s/login", serverContext)
	registerPath := fmt.Sprintf("%s/registration", serverContext)
	k := e.Group(serverContext)

	k.Use(middleware.Logger())
	k.Use(middleware.Recover())
	k.Use(echojwt.WithConfig(echojwt.Config{
		Skipper: func(c echo.Context) bool {
			isSkip := false
			if c.Request().URL.Path == loginPath || c.Request().URL.Path == registerPath {
				isSkip = true
			}
			return isSkip
		},
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte(cfg.JWT.SigningKey),
	}))

	k.POST("/login", authentication.Login)
	k.POST("/registration", authentication.Registration)
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// func restricted(c echo.Context) error {
// 	user := c.Get("user").(*jwt.Token)
// 	claims := user.Claims.(*jwtCustomClaims)
// 	name := claims.Name
// 	return c.String(http.StatusOK, "Welcome "+name+"!")
// }
