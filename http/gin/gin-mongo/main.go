package main

import (
	"github.com/gin-gonic/gin"
	"github.com/williamnoble/goExercises/http/gin/gin-mongo/controllers"
	"github.com/williamnoble/goExercises/http/gin/gin-mongo/internal"
)

func main() {

	db := internal.InitialiseDatabase()
	usersModel := internal.UsersModel{DB: db}
	usersController := controllers.Users{
		UsersModel: usersModel,
	}
	router := getRoutes(usersController)
	_ = router.Run("localhost:9090")

}

func getRoutes(users controllers.Users) *gin.Engine {
	router := gin.Default()
	router.GET("/", users.IndexHandler)
	router.POST("/users", users.CreateUserHandler)
	//router.GET("/users/:id", controllers.GetById)
	//router.GET("/list-users", controllers.AllUser)
	//router.POST("/login", controllers.Login)
	//router.DELETE("/users/:id", controllers.Delete)
	//router.PATCH("/users/:id", controllers.Update)
	return router

}
