package app

import(
	"../controllers"
)

func mapUrls(){
	// Call func in controllers to do the work
	router.GET("/ping", controllers.Ping)

	// Call func in users_controller 
	router.GET("/users", controllers.GetUser)
	router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", controllers.CreateUser)
}