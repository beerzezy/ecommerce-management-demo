package main

import (
	"net/http"

	"github.com/beerzezy/ecommerce-management-demo/cmd/common"
	"github.com/beerzezy/ecommerce-management-demo/cmd/initialization"
	"github.com/beerzezy/ecommerce-management-demo/internal/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	config := initialization.InitConfig()
	db := initialization.InitDB(config)
	defer db.Close()

	e := echo.New()
	//e.Logger = log.Logger()
	e.Validator = &CustomValidator{validator: validator.New()}
	registerRoutes(e, db, config)

	go common.StartServer(e, config)
	common.WaitForGracefulShutdown(e)
}

func registerRoutes(e *echo.Echo, db *gorm.DB, cfg config.Config) {

	// account := account.NewHandler(account.NewService(account.NewRepository(db)))

	// e.GET("/health", health)

	// serverContext := cfg.Server.Context
	// loginPath := fmt.Sprintf("%s/login", serverContext)
	// k := e.Group(serverContext)

	// k.Use(middleware.Logger())
	// k.Use(middleware.Recover())
	// k.Use(echojwt.WithConfig(echojwt.Config{
	// 	Skipper: func(c echo.Context) bool {
	// 		xxx := c.Request().URL.Query()
	// 		fmt.Println("xxxxx,", xxx)
	// 		return c.Request().URL.Path == loginPath
	// 	},
	// 	SigningKey: []byte(cfg.JWT.SigningKey),
	// }))

	// k.POST("/account-info", account.GetAccountInfo)
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
