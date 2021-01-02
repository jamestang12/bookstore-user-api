package app

import (
	"github.com/gin-gonic/gin"
)

var(
	router = gin.Default()
)

func StartApplication(){
	// Call the router
	mapUrls()
	router.Run( ":8080")
}