package routes

import (
	"goapi/controllers"
	"goapi/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisteredRoutes(server *gin.Engine) {
	server.GET("/events", controllers.GetEvents)
	server.GET("/events/:id", controllers.GetEvent)
	server.POST("/events", middlewares.Authenticate, controllers.CreateEvent)
	server.PUT("/events/:id", controllers.UpdateEvent) //PUT is used to Update
	server.DELETE("/events/:id", controllers.DeleteEvent)
	server.POST("/signup", controllers.Signup)
	server.POST("/login", controllers.Login)

}
