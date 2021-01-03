package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../domain/users"
	"fmt"
	"../services"
	"../utils/errors"
)

// func that get call in url_mapings(router)
func GetUser(c *gin.Context){
	c.String(http.StatusNotImplemented, "implement me!")
}

func CreateUser(c *gin.Context){
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invaild json body")
		c.JSON(restErr.Status, restErr)
		fmt.Println(err)
		return 
	}
	
	result, saveErr := services.CreateUser(user)
	if saveErr != nil{
		//c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)

}

func SearchUser(c *gin.Context){
	c.String(http.StatusNotImplemented, "implement me!")
}