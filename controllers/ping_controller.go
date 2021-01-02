package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context){
	// Func to call in url_maping(router)
	c.String(http.StatusOK, "pong")
}