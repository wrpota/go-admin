package zap

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ZapLogger(log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()

		fields := []zapcore.Field{
			zap.String("remote_ip", c.ClientIP()),
			zap.String("latency", time.Since(startTime).String()),
			zap.String("host", c.Request.Host),
			zap.String("request", fmt.Sprintf("%s %s", c.Request.Method, c.Request.RequestURI)),
			zap.Int("status", c.Writer.Status()),
			zap.Int64("size", int64(c.Writer.Size())),
			zap.String("user_agent", c.Request.UserAgent()),
		}
		n := c.Writer.Status()
		switch {
		case n >= 500:
			for _, e := range c.Errors {
				log.With(zap.Error(e)).Error("Server error", fields...)
				return
			}
		case n >= 400:
			for _, e := range c.Errors {
				log.With(zap.Error(e)).Error("Client error", fields...)
				return
			}
		case n >= 300:
			log.Info("Redirection", fields...)
		default:
			log.Info("Success", fields...)
		}
	}
}
