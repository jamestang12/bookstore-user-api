package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"../../bookstore_oauth_go/oauth"
	"../domain/users"
	"../services"
	"../utils/errors"
	"github.com/gin-gonic/gin"
)

// func that get call in url_mapings(router)
func GetUser(c *gin.Context) {

	if err := oauth.AuthenticatRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}

	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if (userErr) != nil {
		err := errors.NewBadRequestError("Invalid user id")
		c.JSON(err.Status, err)
		return
	}
	// Getting user by using the func located in users_service
	user, getErr := services.UsersService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	// If caller id is the same as the user id than it will return a private return which show all the info
	if oauth.GetCallerId(c.Request) == user.Id {
		c.JSON(http.StatusOK, user.Marshall(false))
		return

	}
	// If caller id is not the same as the user id than it will return a public return which show only a part of the info
	c.JSON(http.StatusOK, user.Marshall(oauth.IsPublic(c.Request)))

	//c.String(http.StatusNotImplemented, "implement me!")

}

func CreateUser(c *gin.Context) {
	var user users.User

	// Take in request body than turn JSON to user struct and check if json body is vaild
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invaild json body")
		c.JSON(restErr.Status, restErr)
		fmt.Println(err)
		return
	}

	if callerId := oauth.GetCallerId(c.Request); callerId == 0 {
		err := errors.RestErr{Status: http.StatusUnauthorized, Message: "resource not avai;able"}
		c.JSON(err.Status, err)
	}

	// Create user by using the func located in users_service
	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))

}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func UpdateUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if (userErr) != nil {
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

	result, err := services.UsersService.UpdateUser(isPantial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func DeleteUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if (userErr) != nil {
		err := errors.NewBadRequestError("Invalid user id")
		c.JSON(err.Status, err)
		return
	}

	if err := services.UsersService.DeleUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if (userErr) != nil {
		return 0, errors.NewBadRequestError("Invalid user id")
	}
	return userId, nil
}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.UsersService.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))

}

func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, err := services.UsersService.LoginUser(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))

}
