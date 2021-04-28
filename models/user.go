package models

import (
	mainDb "billing/db/main_db"
	. "billing/entities"
	"log"

	"fmt"
)

// User ...
type UserModel struct {
}

//GetAllUsers Fetch all user data
func (model *UserModel) GetAllUsers() (users []User, err error) {
	if err := mainDb.DB.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

//CreateUser ... Insert New data
func (model *UserModel) CreateUser(user *User) (err error) {
	log.Printf("Received message in user service: %+v\n", user)
	if err = mainDb.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func (model *UserModel) GetUserByID(user *User, id string) (err error) {
	if err = mainDb.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser ... Update user
func (model *UserModel) UpdateUser(user *User, id string) (err error) {
	fmt.Println(user)
	mainDb.DB.Save(user)
	return nil
}

//DeleteUser ... Delete user
func (model *UserModel) DeleteUser(user *User, id string) (err error) {
	mainDb.DB.Where("id = ?", id).Delete(user)
	return nil
}
