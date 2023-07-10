package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/al-kirpichenko/gofermart/internal/models"
	"github.com/al-kirpichenko/gofermart/internal/services/accrual"
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

		Number: string(body),
		Status: "NEW",
		UserID: userID.(uint),
	}

	result := s.DB.Create(&newOrder)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		s.DB.First(&newOrder, "number = ?", newOrder.Number)

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

	// тут будет запрос к сервису начисления баллов
	//TODO запрос к сервесу начисления баллов выполняется асинхронно

	ctx.JSON(http.StatusAccepted, gin.H{"status": "success", "message": "the order has been accepted"})

	go func(order *models.Order, serviceAddress string) {

		var user models.User

		loyalty, err := accrual.GetLoyalty(order.Number, serviceAddress)
		if err != nil {
			log.Println("no response from the accrual service")
			log.Println(err)
			order.Status = "INVALID"
			s.DB.Save(&order)
			return
		}

		order.Accrual = loyalty.Accrual
		order.Status = loyalty.Status

		result := s.DB.First(&user, "id = ?", userID)
		if result.Error != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "User not found"})
			return
		}
		user.Balance = user.Balance + order.Accrual

		s.DB.Save(&order)
		s.DB.Save(&user)

	}(&newOrder, s.config.ServiceAddress)

}

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
