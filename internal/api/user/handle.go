package user

import "github.com/gin-gonic/gin"

var _ Handler = (*handler)(nil)

type Handler interface {
	i()
	HelloWorld() gin.HandlerFunc
}

type handler struct {
}

func (h *handler) i() {}

//返回handler 对象
func New() Handler {
	return &handler{}
}
