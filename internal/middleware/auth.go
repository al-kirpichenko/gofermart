package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/al-kirpichenko/gofermart/internal/services/jwt"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userCookie, err := ctx.Cookie("token")

		if err != nil {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userID, err := jwt.GetUserIDFromToken(userCookie)

		if err != nil || userID == 0 {

			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("userID", userID)

		ctx.Next()

	}
}
