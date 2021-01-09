package app

import (
	"github.com/gin-gonic/gin"
	"../logger"
)

var(
	router = gin.Default()
)

func StartApplication(){
	// Call the router
	mapUrls()
	logger.Info("About to start the application.....")
	router.Run( ":8080")
}