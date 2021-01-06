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


func UpdateUser(c *gin.Context)  {
	userId, userErr := strconv.ParseInt(c.Param("user_id"),10, 64)
	if(userErr) != nil{
		err := errors.NewBadRequestError("Invalid user id")
		c.JSON(err.Status, err)
		return
	}
	var user users.User
	
	// Take in request body than turn JSON to user struct and check if json body is vaild
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invaild json body")
		c.JSON(restErr.Status, restErr)
		fmt.Println(err)
		return 
	}

	user.Id = userId

	isPantial := c.Request.Method == http.MethodPatch


	result, err := services.UpdateUser(isPantial,user)
	if err != nil{
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func DeleteUser(c *gin.Context){
	userId, userErr := strconv.ParseInt(c.Param("user_id"),10, 64)
	if(userErr) != nil{
		err := errors.NewBadRequestError("Invalid user id")
		c.JSON(err.Status, err)
		return
	}

	if err := services.DeleUser(userId); err != nil{
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func getUserId(userIdParam string)(int64, *errors.RestErr){
	userId, userErr := strconv.ParseInt(userIdParam,10, 64)
	if(userErr) != nil{
		return 0, errors.NewBadRequestError("Invalid user id")
	}
	return userId,nil
}

func Search(c *gin.Context){
	status := c.Query("status")

	users, err := services.Search(status)
	if err != nil{
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users)

}