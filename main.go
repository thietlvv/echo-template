package main

import (
	"log"
	"os"

	"billing/serv/urls"

	"billing/db/main_db"
	middlewares "billing/middlewares"
	"billing/migrations"
	l "billing/utils/logger"
	"billing/utils/pubsub/nats"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func initRouterDefault() *gin.Engine {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	r := gin.Default()

	// Use global middleware
	r.Use(middlewares.RequestLog())

	// Recovery middleware recovers from any panics and writes a 500 if there was one
	r.Use(gin.Recovery())
	// r.Use(sentry.Recovery(raven.DefaultClient, false))

	return r
}

// Initiate web server
func main() {

	// Load the .env file in the current directory
	// .env already load in-case run by docker-compose
	err := godotenv.Load()
	if err != nil {
		l.Logger("", "").Warningln("Cannot get config from .env file manual - if run from docker-compose skip this warning", err)
	} else {
		l.Logger("", "").Infoln("Get config from .env successful")
	}

	// Init Log
	l.InitLogger()

	// Start DB
	db, err := main_db.InitDB()
	if err != nil {
		l.Logger("", "").Warningln("Can't connect RDB. ", err)
	} else {
		l.Logger("", "").Infoln("Database connect successful")
		defer db.Close()

		// Auto migration
		migration := os.Getenv("RDB_AUTO_MIGRATION")
		if migration == "1" {
			migrations.MigrateDataTable()
		}
	}

	err = nats.InitNats()
	if err != nil {
		l.Logger("", "").Warningln("could not connect to NATS server: ", err)
	} else {
		l.Logger("", "").Infoln("NATS server connect successful")
	}

	// Init routes default
	r := initRouterDefault()

	// Init url routes
	urls.InitUrlsBilling(r)

	// Start server
	runport := os.Getenv("RUN_PORT")
	if runport == "" {
		runport = "9090"
	}

	log.Fatal(r.Run(":" + runport)) // running
}
