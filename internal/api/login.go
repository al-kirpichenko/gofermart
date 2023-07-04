package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/al-kirpichenko/gofermart/internal/models"
	"github.com/al-kirpichenko/gofermart/internal/services"
)

// аутентификация пользователя

func (s *Server) Login(ctx *gin.Context) {

	var auth *models.Auth
	var user models.User

	if err := ctx.ShouldBindJSON(&auth); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	result := s.DB.First(&user, "login = ?", auth.Login)

	if result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Invalid Login or Password"})
		return
	}

	if err := services.VerifyPassword(user.Password, auth.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Invalid Login or Password"})
		return
	}

}
