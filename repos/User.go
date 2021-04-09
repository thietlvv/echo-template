//Models/User.go

package repos

import (
	db "billing/db/main_db"
	"billing/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllUsers Fetch all user data
func GetAllUsers(user *[]models.User) (err error) {
	if err = db.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

//CreateUser ... Insert New data
func CreateUser(user *models.User) (err error) {
	if err = db.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetUserByID(user *models.User, id string) (err error) {
	if err = db.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser ... Update user
func UpdateUser(user *models.User, id string) (err error) {
	fmt.Println(user)
	db.DB.Save(user)
	return nil
}

//DeleteUser ... Delete user
func DeleteUser(user *models.User, id string) (err error) {
	db.DB.Where("id = ?", id).Delete(user)
	return nil
}
