package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/al-kirpichenko/gofermart/internal/models"
)

// получение текущего баланса счёта баллов лояльности пользователя

func (s *Server) Balance(ctx *gin.Context) {

	var (
		balance models.Balance
		user    models.User
	)

	userID, _ := ctx.Get("userID")

	result := s.DB.First(&user, "id = ?", userID)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "User not found"})
		return
	}

	balance.Current = user.Balance
	balance.Withdrawn = user.Withdrawn

	response, err := json.Marshal(balance)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "json error"})
		return
	}

	ctx.Data(http.StatusOK, "application/json", response)

}
