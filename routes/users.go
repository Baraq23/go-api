package routes

import (
	"goapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func signUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully."})
}