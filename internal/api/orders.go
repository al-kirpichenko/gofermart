package api

import (
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/al-kirpichenko/gofermart/internal/models"
	"github.com/al-kirpichenko/gofermart/internal/services/luhn"
)

// загрузка пользователем номера заказа для расчёта

func (s *Server) AddOrder(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	num, err := strconv.Atoi(string(body))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid value", "message": err.Error()})
		return
	}

	if !luhn.Valid(num) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"status": "fail", "message": "Invalid number(Luhn)"})
		return
	}

	userID, _ := ctx.Get("userID")

	newOrder := models.Order{

		Number: num,
		Status: "NEW",
		UserID: userID.(uint),
	}

	result := s.DB.Create(&newOrder)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		s.DB.First(&newOrder)

		if newOrder.UserID == userID.(uint) {
			ctx.JSON(http.StatusOK, gin.H{"status": "fail", "message": "This Order has been loaded this user"})
			return
		} else {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "This Order has been loaded other user"})
			return
		}
	} else if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Something bad happened"})
		return
	}

}

// получение списка загруженных пользователем номеров заказов, статусов их обработки и информации о начислениях

func (s *Server) GetOrders(c *gin.Context) {

}

// получение информации о расчёте начислений баллов лояльности

func (s *Server) Order(c *gin.Context) {

}
