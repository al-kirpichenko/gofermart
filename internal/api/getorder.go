package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/al-kirpichenko/gofermart/internal/models"
)

// получение списка загруженных пользователем номеров заказов, статусов их обработки и информации о начислениях

func (s *Server) GetOrders(ctx *gin.Context) {

	var orders []models.Order

	userID, _ := ctx.Get("userID")

	s.DB.Order("created_at").Where("user_id = ?", userID).Find(&orders)

	if len(orders) == 0 {
		ctx.JSON(http.StatusNoContent, gin.H{"status": "success", "message": "No content"})
		return
	}

	response, err := json.Marshal(orders)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "json error"})
		return
	}

	ctx.Data(http.StatusOK, "application/json", response)
}
