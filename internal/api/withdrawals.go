package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/al-kirpichenko/gofermart/internal/models"
)

// получение информации о выводе средств с накопительного счёта пользователем

func (s *Server) Withdrawals(ctx *gin.Context) {

	var (
		user        models.User
		withdrawals []models.Withdraw
	)

	userID, _ := ctx.Get("userID")

	result := s.DB.First(&user, "id = ?", userID)

	if result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "User not found"})
		return
	}

	s.DB.Order("created_at").Where("user_id = ?", userID).Find(&withdrawals)

	if len(withdrawals) == 0 {
		ctx.JSON(http.StatusNoContent, gin.H{"status": "success", "message": "No content"})
		return
	}

	response, err := json.Marshal(withdrawals)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "json error"})
		return
	}

	ctx.Data(http.StatusOK, "application/json", response)

}
