package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../domain/users"
	"fmt"
	"../services"
	"../utils/errors"
	"strconv"
)

// func that get call in url_mapings(router)
func GetUser(c *gin.Context){
	userId, userErr := strconv.ParseInt(c.Param("user_id"),10, 64)
	if(userErr) != nil{
		err := errors.NewBadRequestError("Invalid user id")
		c.JSON(err.Status, err)
		return
	}
	// Getting user by using the func located in users_service
	user, getErr := services.GetUser(userId)
	if getErr != nil{
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)

	//c.String(http.StatusNotImplemented, "implement me!")


}

func CreateUser(c *gin.Context){
	var user users.User
	
	// Take in request body than turn JSON to user struct and check if json body is vaild
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invaild json body")
		c.JSON(restErr.Status, restErr)
		fmt.Println(err)
		return 
	}
	
	// Create user by using the func located in users_service
	result, saveErr := services.CreateUser(user)
	if saveErr != nil{
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)



}

func SearchUser(c *gin.Context){
	c.String(http.StatusNotImplemented, "implement me!")
}