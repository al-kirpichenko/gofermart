package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/al-kirpichenko/gofermart/internal/models"
	"github.com/al-kirpichenko/gofermart/internal/services/luhn"
	"github.com/al-kirpichenko/gofermart/internal/services/math"
)

// запрос на списание баллов с накопительного счёта в счёт оплаты нового заказа

func (s *Server) Withdraw(ctx *gin.Context) {

	var (
		user     models.User
		withdraw models.Withdraw
	)

	if err := ctx.ShouldBindJSON(&withdraw); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	num, err := strconv.Atoi(withdraw.Order)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid order", "message": err.Error()})
		return
	}

	if !luhn.Valid(num) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"status": "fail", "message": "Invalid number(Luhn)"})
		return
	}

	s.DB.Transaction(func(tx *gorm.DB) error {

		userID, _ := ctx.Get("userID")

		result := tx.First(&user, "id = ?", userID)

		if result.Error != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "User not found"})
			return err
		}

		if user.Balance-withdraw.Sum < 0 {
			ctx.JSON(http.StatusPaymentRequired, gin.H{"status": "fail", "message": "need more gold..."})
			return err
		}

		user.Balance = math.RoundFloat(user.Balance-withdraw.Sum, 2)
		user.Withdrawn = user.Withdrawn + withdraw.Sum
		withdraw.UserID = user.ID
		withdraw.ProcessedAt = time.Now()

		r := tx.Save(&withdraw)

		if r.Error != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": r.Error})
			s.Logger.Error("Don't save withdraw! ", zap.Error(r.Error))
			return err
		}

		if err := tx.Save(&user).Error; err != nil {
			return err
		}
		//s.DB.Save(&user)
		return nil
	})
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "the withdraw has been accepted"})

}
