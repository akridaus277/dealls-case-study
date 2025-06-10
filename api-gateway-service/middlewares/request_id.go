package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestIDKey = "X-Request-ID"

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.GetHeader(RequestIDKey)
		if reqID == "" {
			reqID = uuid.NewString()
		}

		// Set ke header supaya diteruskan ke downstream service
		c.Request.Header.Set(RequestIDKey, reqID)
		c.Writer.Header().Set(RequestIDKey, reqID)

		// Simpan ke Gin Context
		c.Set(RequestIDKey, reqID)

		c.Next()
	}
}
