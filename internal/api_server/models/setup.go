package models

import (
	"fmt"
	"log"
	"net/url"

	"github.com/BulatKarimov/gin_rest_api/internal/api_server/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {

	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal(err)
		return
	}

	dsn := url.URL{
		User:     url.UserPassword(config.DbUser, config.DbPassword),
		Scheme:   config.DbDriver,
		Host:     fmt.Sprintf("%s:%d", config.DbHost, config.DbPort),
		Path:     "gin_rest_api",
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}

	database, err := gorm.Open(config.DbDriver, dsn.String())

	if err != nil {
		log.Fatal(err)
	}

	// database.AutoMigrate(&Event{})

	DB = database
}
