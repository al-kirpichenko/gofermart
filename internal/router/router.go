package router

import (
	"github.com/gin-gonic/gin"

	"github.com/al-kirpichenko/gofermart/internal/api"
)

func Router(server *api.Server) *gin.Engine {

	r := gin.Default()

	r.POST("/api/user/register", server.Register)
	r.POST("/api/user/login", server.Login)
	r.POST("/api/user/orders", server.AddOrder)
	r.GET("/api/user/orders", server.GetOrders)
	r.GET("/api/user/orders/:number", server.Order)
	r.GET("/api/user/balance", server.Balance)
	r.POST("/api/user/balance/withdraw", server.Withdraw)
	r.POST("/api/user/withdrawals", server.Withdrawals)

	return r
}
