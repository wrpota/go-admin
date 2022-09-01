package main

import (
	"github.com/wrpota/go-gin/configs"
	_ "github.com/wrpota/go-gin/init"
	"github.com/wrpota/go-gin/internal/router"
)

func main() {
	s := router.NewHttpServer()

	s.Run("localhost:" + configs.Get().GetString("HttpServer.Web.Port"))
}
