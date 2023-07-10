package router

import (
	"github.com/gin-contrib/gzip"

	"github.com/gin-gonic/gin"

	"github.com/al-kirpichenko/gofermart/internal/api"
	"github.com/al-kirpichenko/gofermart/internal/middleware"
)

func Router(server *api.Server) *gin.Engine {

	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	auth := r.Group("/")

	{
		auth.Use(middleware.Auth())
		auth.POST("/api/user/orders", server.AddOrder)
		auth.GET("/api/user/orders", server.GetOrders)
		auth.GET("/api/user/balance", server.Balance)
		auth.POST("/api/user/balance/withdraw", server.Withdraw)
		auth.GET("/api/user/withdrawals", server.Withdrawals)
	}

	r.POST("/api/user/register", server.Register)
	r.POST("/api/user/login", server.Login)

	return r
}
