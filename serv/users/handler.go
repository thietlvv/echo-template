package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	. "billing/entities"
	"billing/helper"
	"billing/models"
	"billing/utils/pubsub/nats"
	"billing/utils/response"

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
		response.Success(c, "", user)
	}
}

//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var userModel models.UserModel
	var user User
	c.ShouldBindJSON(&user)
	err := userModel.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response.Success(c, "", &user)
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
		response.Success(c, "", user)
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
		response.Success(c, "", user)
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

// PublishUserCreated publish an event via NATS server
func PublishUserCreated(c *gin.Context) {

	// Publish order to NATS server
	user := User{
		Name:  "ggggggggggg",
		Email: "hh@gmail.com",
	}
	userData, err := json.Marshal(user)
	if err != nil {
		log.Printf("could not marshal user: %v\n", err)
		return
	}

	// Publish message on subject
	err = nats.Pub(Subject, userData)
	if err != nil {
		log.Printf("could not publish user: %v\n", err)
		return
	}
	c.JSON(http.StatusOK, event)
}
