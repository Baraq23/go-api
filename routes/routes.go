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
	authenticated.GET("/events/registered", controllers.GetRegisteredEvents)
	authenticated.GET("/events/created", controllers.GetCreatedEvents)
	authenticated.PUT("/events/:id", controllers.UpdateEvent) //PUT is used to Update
	authenticated.DELETE("/events/:id", controllers.DeleteEvent)
	authenticated.POST("/events/:id/register", controllers.RegisterForEvent)
	authenticated.DELETE("/events/:id/register", controllers.CancelRegistration)

	server.POST("/signup", controllers.Signup)
	server.POST("/login", controllers.Login)

}
