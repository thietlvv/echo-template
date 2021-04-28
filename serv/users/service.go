package users

import (
	"billing/entities"
	"billing/models"
	l "billing/utils/logger"
	"billing/utils/pubsub/nats"
	"encoding/json"
	"log"
)

func HandleCreateOrderSub() error {
	var userModel models.UserModel
	user := entities.User{}

	err := nats.Sub(Subject, func(data []byte) {
		var err error
		err = json.Unmarshal(data, &user)
		if err != nil {
			l.Logger("", "").Errorln("could not unmarshal user: ", err)
			return
		}
		// Handle the message

		log.Printf("&user:: %+v\n", user)

		err = userModel.CreateUser(&user)
		if err != nil {
			l.Logger("", "").Errorln("could not store users: ", err)
			return
		}
		l.Logger("", "").Infoln("store users successfully")
	})

	if err != nil {
		l.Logger("", "").Errorln("could not subscribe to users: ", err)
		return err
	}

	return nil
}
