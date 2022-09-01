package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wrpota/go-gin/internal/api/auth"
	"github.com/wrpota/go-gin/internal/api/user"
	"github.com/wrpota/go-gin/internal/middleware"
)

func setApiRouter(r *gin.Engine) {
	api := r.Group("/api")

	// 登录
	var authRouter = api.Group("auth")
	authHandle := auth.New()
	authRouter.POST("/login", authHandle.Login())
	authRouter.Use(middleware.JWTAuth()).PUT("/refreshToken", authHandle.RefreshToken())

	var userRouter = api.Group("/user")
	user := user.New()
	userRouter.GET("", user.HelloWorld())
}
