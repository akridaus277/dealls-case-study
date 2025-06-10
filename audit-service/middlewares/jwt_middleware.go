package middlewares

import (
	"audit-service/utils"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			utils.SendResponse(c, utils.NewBadRequestResponse("Missing or invalid token"))
			return
		}
		log.Printf("authheader : %s", authHeader)

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			utils.SendResponse(c, utils.NewBadRequestResponse("Invalid token"))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.SendResponse(c, utils.NewBadRequestResponse("Invalid claims"))
			return
		}

		c.Set("user", claims["user"])
		c.Set("role", claims["role"])
		c.Next()
	}
}
