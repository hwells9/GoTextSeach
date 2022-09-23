package controllers

import (
	"bench/textsearch/authentication"
	"bench/textsearch/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register new user
func RegisterUser(context *gin.Context) {
	var user authentication.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := database.Db.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}
