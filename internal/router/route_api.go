package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wrpota/go-gin/internal/api/user"
)

func setApiRouter(r *gin.Engine) {
	api := r.Group("/api")
	var userRouter = api.Group("/user")
	user := user.New()
	userRouter.GET("", user.HelloWorld())
}
