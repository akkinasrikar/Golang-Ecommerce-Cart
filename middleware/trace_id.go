package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TraceIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := uuid.NewString()
		c.Header("X-Trace-Id", traceID)
		c.Set("traceID", traceID)
		c.Next()
	}
}
