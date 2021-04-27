package users

import (
	"fmt"
	"net/http"

	. "billing/entities"
	"billing/helper"
	"billing/models"

	"github.com/gin-gonic/gin"
)

//GetUsers ... Get all users
func GetUsers(c *gin.Context) {
	var userModel models.UserModel
	requestID := c.GetString("x-request-id")
	helper.Logger(requestID, "").Infoln("RequestID= ", requestID)
	// cacheTest := helper.CacheExists("xxxxxxxxxx")
	// helper.Logger(requestID, "").Infoln("cacheTest= ", cacheTest)

	// httpCode, body, erro := helper.MakeHTTPRequest("GET", "https://api-101.glitch.me/customers", "", nil, true)
	// helper.Logger(requestID, "").Infoln("httpCode= ", httpCode)
	// helper.Logger(requestID, "").Infoln("body= ", fmt.Sprintf("%s", body))
	// helper.Logger(requestID, "").Infoln("error= ", erro)

	var user []User
	user, err := userModel.GetAllUsers()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var userModel models.UserModel
	var user User
	c.BindJSON(&user)
	err := userModel.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	var userModel models.UserModel

	id := c.Params.ByName("id")
	var user User
	err := userModel.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//UpdateUser ... Update the user information
func UpdateUser(c *gin.Context) {
	var userModel models.UserModel

	var user User
	id := c.Params.ByName("id")
	err := userModel.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = userModel.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	var userModel models.UserModel

	var user User
	id := c.Params.ByName("id")
	err := userModel.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
