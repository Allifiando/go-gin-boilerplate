package server

import (
	"io"
	"os"

	"github.com/Allifiando/go-gin-boilerplate/controller"
	"github.com/Allifiando/go-gin-boilerplate/middleware"
	"github.com/gin-gonic/gin"
)

// var apiVersion string = "/v1"

// Init ...
func Init() *gin.Engine {
	r := gin.Default()
	r.RedirectTrailingSlash = true
	r.RedirectFixedPath = true
	r.RemoveExtraSlash = true

	// add loging file
	f, _ := os.Create("log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r.Use(middleware.CORSMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "OK"})
	})

	users := r.Group("/user")
	userController := controller.User{}
	users.POST("/login", userController.Login)
	users.POST("/register", userController.Register)
	users.Use(middleware.Auth())
	{
		users.GET("/", userController.ListUser)
		users.GET("/uuid/:uuid", userController.FindOneByUUID)
	}

	// glosarium := r.Group("/glosarium")
	// glosariumController := controllers.Glosarium{}
	// {
	// 	glosarium.GET("/", glosariumController.ListAll)
	// 	glosarium.POST("/", glosariumController.Create)
	// 	glosarium.DELETE("/id/:id", glosariumController.Delete)
	// }

	return r
}
