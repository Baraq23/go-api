package routes

import (
	"goapi/controllers"

	"github.com/gin-gonic/gin"
)

func RegisteredRoutes(server *gin.Engine) {
	server.GET("/events", controllers.GetEvents)
	server.GET("/events/:id", controllers.GetEvent)
	server.POST("/events", controllers.CreateEvent)
	server.PUT("/events/:id", controllers.UpdateEvent) //PUT is used to Update
	server.DELETE("/events/:id", controllers.DeleteEvent)
	server.POST("/signup", controllers.SignUp)

}
