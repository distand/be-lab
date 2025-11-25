package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Log() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - %s [%s %s] [%d] [%s] %s \"%s\" %s\n",
			param.TimeStamp.Format(time.DateTime),
			param.ClientIP,
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
			param.Request.Proto,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}
