package main

import (
	"billing/db/main_db"
	"billing/migrations"
	"billing/serv/users"
	"billing/utils/logger"
	"billing/utils/pubsub/nats"
	"os"
	"runtime"

	"github.com/joho/godotenv"
)

func main() {

	// Load the .env file in the current directory
	// .env already load in-case run by docker-compose
	err := godotenv.Load()
	if err != nil {
		logger.Logger("", "").Warningln("Cannot get config from .env file manual - if run from docker-compose skip this warning", err)
	} else {
		logger.Logger("", "").Infoln("Get config from .env successful")
	}

	// Init Log
	logger.InitLogger()

	// Start DB
	_, err = main_db.InitDB()
	if err != nil {
		logger.Logger("", "").Warningln("Can't connect RDB. ", err)
	} else {
		logger.Logger("", "").Infoln("Database connect successful")
		// defer db.Close()

		// Auto migration
		migration := os.Getenv("RDB_AUTO_MIGRATION")
		if migration == "1" {
			migrations.MigrateDataTable()
		}
	}

	err = nats.InitNats()
	if err != nil {
		logger.Logger("", "").Warningln("could not connect to NATS server: ", err)
	} else {
		logger.Logger("", "").Infoln("NATS server connect successful")
	}

	err = users.HandleCreateOrderSub()
	if err != nil {
		logger.Logger("", "").Errorln("[HandleCreateOrderSub] Error: ", err)
	}

	// err = users.HandleCreateOrderQueueSub()
	// if err != nil {
	// 	logger.Logger("", "").Errorln("[HandleCreateOrderQueueSub-1]Error: ", err)
	// }

	// err = users.HandleCreateOrderQueueSub()
	// if err != nil {
	// 	logger.Logger("", "").Errorln("[HandleCreateOrderQueueSub-2]Error: ", err)
	// }

	// Keep the connection alive
	runtime.Goexit()
}
