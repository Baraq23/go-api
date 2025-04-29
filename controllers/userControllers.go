package controllers

import (
	"goapi/models"
	"goapi/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)
var TotalUsers int

func Signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	
	// if user.Name == "" {
	// 	user.Name = "User"
	// }

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}

	TotalUsers+=1

	context.JSON(http.StatusCreated, gin.H{"message": "User signed up successfully.", "total users": TotalUsers})
}


func Login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.ValidateCridentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate user token."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User logged in successfully.", "token": token, "userId": user.ID})
}


func GetCreatedEvents(context *gin.Context) {
	userId := context.GetInt64("userId")
	events, err := models.GetCreatedEvent(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events created. Try agin later."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"list of your created events": events})
}

func GetRegisteredEvents(context *gin.Context) {
	userId := context.GetInt64("userId")
	events, err := models.GetRegisteredEvent(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events registered to. Try agin later."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"list of your registered events": events})
}

