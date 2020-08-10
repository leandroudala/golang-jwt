package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// PORT that apps will run
	PORT = 0
	// DBURL is db connection string
	DBURL = ""
	// DBDRIVER is the driver
	DBDRIVER = ""
)

// Load a config
func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Println(err)
		PORT = 3000
	}
	log.Println("This app will run on port", PORT)

	DBDRIVER = os.Getenv("DB_DRIVER")
	DBURL = fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
}
