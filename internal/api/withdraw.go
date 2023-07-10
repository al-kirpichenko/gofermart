package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/al-kirpichenko/gofermart/internal/models"
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

	res := s.DB.First(&withdraw, "order = ?", withdraw.Order)

	if res.Error != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"status": "fail", "message": "Order not found"})
		return
	}

	userID, _ := ctx.Get("userID")

	result := s.DB.First(&user, "id = ?", userID)

	if result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "User not found"})
		return
	}

	if user.Balance-withdraw.Sum < 0 {
		ctx.JSON(http.StatusPaymentRequired, gin.H{"status": "fail", "message": "need more gold..."})
		return
	}

	user.Balance = user.Balance - withdraw.Sum

	s.DB.Save(&withdraw)
	s.DB.Save(&user)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "the withdraw has been accepted"})

}
