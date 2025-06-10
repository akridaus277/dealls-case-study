package middlewares

import (
	"fmt"
	"log"
	"os"
	"overtime-service/utils"
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

func JWTAuthAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			utils.SendResponse(c, utils.NewBadRequestResponse("Missing or invalid token"))
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			utils.SendResponse(c, utils.NewBadRequestResponse("Invalid token"))
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.SendResponse(c, utils.NewBadRequestResponse("Invalid claims"))
			c.Abort()
			return
		}

		rolesClaim, ok := claims["roles"]
		if !ok {
			utils.SendResponse(c, utils.NewBadRequestResponse("Roles not found in token"))
			c.Abort()
			return
		}

		// Pastikan roles adalah slice atau string yang bisa di-split
		var roles []string
		switch v := rolesClaim.(type) {
		case []interface{}:
			for _, r := range v {
				roles = append(roles, fmt.Sprintf("%v", r))
			}
		case string:
			roles = strings.Split(v, ",")
		default:
			utils.SendResponse(c, utils.NewBadRequestResponse("Invalid roles format"))
			c.Abort()
			return
		}

		log.Printf("Roles in token: %v", roles)

		// Cek apakah ada role "admin"
		hasAdmin := false
		for _, r := range roles {
			if strings.Contains(strings.TrimSpace(r), "ADMIN") {
				hasAdmin = true
				break
			}
		}

		if !hasAdmin {
			utils.SendResponse(c, utils.NewUnauthorizedResponse("Access denied: Admin only"))
			c.Abort()
			return
		}

		c.Set("user", claims["user"])
		c.Set("roles", roles)
		c.Next()
	}
}
