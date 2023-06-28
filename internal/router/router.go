package router

import (
	"github.com/gin-gonic/gin"

	"github.com/al-kirpichenko/gofermart/internal/app"
)

func Router(app *app.Application) *gin.Engine {

	r := gin.Default()

	r.POST("/api/user/register", app.Register)
	r.POST("/api/user/login", app.Login)
	r.POST("/api/user/orders", app.AddOrder)
	r.GET("/api/user/orders", app.GetOrders)
	r.GET("/api/user/balance", app.Balance)
	r.POST("/api/user/balance/withdraw", app.Withdraw)
	r.POST("/api/user/withdrawals", app.Withdrawals)

	return r
}
