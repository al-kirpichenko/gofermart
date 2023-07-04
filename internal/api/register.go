package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/al-kirpichenko/gofermart/internal/models"
	"github.com/al-kirpichenko/gofermart/internal/services"
)

// регистрация пользователя

func (s *Server) Register(ctx *gin.Context) {

	var auth *models.Auth

	if err := ctx.ShouldBindJSON(&auth); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	hashedPassword, err := services.HashPassword(auth.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	newUser := models.User{

		Login:    auth.Login,
		Password: hashedPassword,
	}

	result := s.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User with that email already exists"})
		return
	} else if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Something bad happened"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "the user is registered"})

}
