package routes

import "github.com/gin-gonic/gin"

func RegisteredRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent) //PUT is used to Update
	server.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signUp)

}
