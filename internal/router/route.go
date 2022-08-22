package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wrpota/go-gin/configs"
	"github.com/wrpota/go-gin/internal/global/variable"
	"github.com/wrpota/go-gin/internal/middleware"
	"github.com/wrpota/go-gin/internal/pkg/zap"
	"github.com/wrpota/go-gin/pkg/env"
	zaplog "go.uber.org/zap"
)

func NewHttpServer() *gin.Engine {
	if env.Active().IsPro() {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	//根据配置进行设置跨域
	if configs.Get().GetBool("HttpServer.AllowCrossDomain") {
		router.Use(middleware.CORS())
	}

	// 设置可信任的代理服务器列表
	if configs.Get().GetBool("HttpServer.TrustProxies.IsOpen") {
		if err := router.SetTrustedProxies(configs.Get().GetStringSlice("HttpServer.TrustProxies.ProxyServerList")); err != nil {
			variable.ZapLog.Error("Gin 设置信任代理服务器出错", zaplog.Error(err))
		}
	} else {
		_ = router.SetTrustedProxies(nil)
	}

	gin.DisableConsoleColor()

	router.Use(zap.ZapLogger(variable.GinZapLog))

	setApiRouter(router)

	return router
}
