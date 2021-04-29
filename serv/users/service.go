package users

import (
	"billing/entities"
	"billing/models"
	l "billing/utils/logger"
	"billing/utils/pubsub/nats"
	"encoding/json"
)

func HandleCreateOrderSub() error {
	err := nats.Sub(Subject, func(data []byte) {
		var userModel *models.UserModel

		var err error
		var user *entities.User
		err = json.Unmarshal(data, &user)
		if err != nil {
			l.Logger("", "").Errorln("could not unmarshal user: ", err)
			return
		}
		// Handle the message
		err = userModel.CreateUser(user)
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

func HandleCreateOrderQueueSub() error {
	err := nats.QueueSub(Subject, Queue, func(data []byte) {
		var userModel *models.UserModel

		var err error
		var user *entities.User
		err = json.Unmarshal(data, &user)
		if err != nil {
			l.Logger("", "").Errorln("could not unmarshal user: ", err)
			return
		}
		// Handle the message
		err = userModel.CreateUser(user)
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
