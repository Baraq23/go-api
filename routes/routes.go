package routes

import (
	"goapi/controllers"
	"goapi/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisteredRoutes(server *gin.Engine) {
	server.GET("/events", controllers.GetEvents)
	server.GET("/events/:id", controllers.GetEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate,)
	authenticated.POST("/events", controllers.CreateEvent)
	authenticated.PUT("/events/:id", controllers.UpdateEvent) //PUT is used to Update
	authenticated.DELETE("/events/:id", controllers.DeleteEvent)
	
	server.POST("/signup", controllers.Signup)
	server.POST("/login", controllers.Login)

}
