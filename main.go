package main

import (
	"log"
	"net/http"
	"os"

	"billing/serv/urls"

	db "billing/db/main_db"
	l "billing/utils/logger"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var err error

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// Initiate web server
func main() {

	// Load the .env file in the current directory
	// .env already load in-case run by docker-compose
	err = godotenv.Load()
	if err != nil {
		l.Logger("", "").Warningln("Cannot get config from .env file manual - if run from docker-compose skip this warning", err)
	} else {
		l.Logger("", "").Infoln("Get config from .env successful")
	}

	// Init Log
	l.InitLogger()

	// Start DB
	err = db.InitDB()
	if err != nil {
		l.Logger("", "").Warningln("Can't connect RDB. ", err)
	} else {
		l.Logger("", "").Infoln("Database connect successful")
		defer db.DB.Close()

		// Auto migration
		migration := os.Getenv("RDB_AUTO_MIGRATION")
		if migration == "1" {
			db.MigrateDataTable()
		}
	}

	e := echo.New()
	e.Use(middleware.Logger())

	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodDelete, http.MethodPut},
		MaxAge:       86400,
	}))
	e.GET("/", hello)

	// Initial Urls
	urls.InitUrlsBilling(e)

	// Start server
	runport := os.Getenv("RUN_PORT")
	if runport == "" {
		runport = "9090"
	}

	log.Fatal(e.Start(":" + runport))
}
