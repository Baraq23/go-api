package controllers

import (
	"goapi/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterForEvent(context *gin.Context) {

	userId := context.GetInt64("userId")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event egistration successful."})
}

func CancelRegistration(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	_, err = models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	userId := context.GetInt64("userId")
	err = models.Deregister(userId, eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not deregister from the event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "You have derigistered from the event successfully"})

}
