package initialization

import (
	"fmt"
	"strings"
	"time"

	"github.com/beerzezy/ecommerce-management-demo/internal/config"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func InitConfig() config.Config {
	var c config.Config
	viper.SetDefault("database.sslmode", "disable")
	viper.SetConfigName("config")
	viper.AddConfigPath("./config/")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("config file not found")
		} else {
			panic(fmt.Errorf("fatal error config file: %s", err))
		}
	}
	viper.AutomaticEnv()
	viper.Unmarshal(&c)
	return c
}

func InitDB(c config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.DBName,
		c.Database.Username,
		c.Database.Password,
		c.Database.SSLMode,
	)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(fmt.Errorf("error connecting to DB: %v", err))
	}
	db.LogMode(true)
	//db.SetLogger(log.GORMLogger{})
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 300)
	if err != nil {
		panic(fmt.Errorf("error connecting to DB: %v", err))
	}

	return db
}
