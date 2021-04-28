package main

import (
	. "billing/serv/users"
	l "billing/utils/logger"
	"billing/utils/pubsub/nats"
	"runtime"
)

func main() {
	err := nats.InitNats()
	if err != nil {
		l.Logger("", "").Warningln("could not connect to NATS server: ", err)
	} else {
		l.Logger("", "").Infoln("NATS server connect successful")
	}

	_ = HandleCreateOrderSub()

	// Keep the connection alive
	runtime.Goexit()
}
