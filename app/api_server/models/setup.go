package models

import (
	"fmt"
	"log"
	"net/url"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := url.URL{
		User:     url.UserPassword("postgres", "112233"),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", "localhost", 5432),
		Path:     "gin_rest_api",
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}

	database, err := gorm.Open("postgres", dsn.String())

	if err != nil {
		log.Fatal(err)
	}

	// database.AutoMigrate(&Event{})

	DB = database
}
