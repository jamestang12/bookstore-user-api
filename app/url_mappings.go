package app

import(
	"../controllers"
)

func mapUrls(){
	// Call func in controllers to do the work
	router.GET("/ping", controllers.Ping)

	// Call func in users_controller 
	router.GET("/users/:user_id", controllers.GetUser)
	// router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", controllers.CreateUser)
	router.PUT("/users/:user_id", controllers.UpdateUser)
	router.PATCH("/users/:user_id", controllers.UpdateUser)
	router.DELETE("/users/:user_id", controllers.DeleteUser)


}