package api

import (
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"

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
		}
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "This Order has been loaded other user"})
		return

	} else if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Something bad happened"})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"status": "success", "message": "the order has been accepted"})

	var user models.User

	loyalty, err := accrual.GetLoyalty(newOrder.Number, s.config.ServiceAddress)

	if err != nil {
		s.Logger.Error("No response from the accrual service", zap.Error(err))
		newOrder.Status = "PROCESSING"
		s.DB.Save(&newOrder)
		return
	}

	s.DB.Transaction(func(tx *gorm.DB) error {

		newOrder.Accrual = loyalty.Accrual
		newOrder.Status = loyalty.Status

		res := s.DB.First(&user, "id = ?", userID)

		if res.Error != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "User not found"})
			return err
		}
		user.Balance = user.Balance + newOrder.Accrual

		if err := tx.Save(&newOrder).Error; err != nil {
			return err
		}
		if err := tx.Save(&user).Error; err != nil {
			return err
		}
		//s.DB.Save(&newOrder)
		//s.DB.Save(&user)
		return nil
	})

}
