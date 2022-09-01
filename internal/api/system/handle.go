package system

var _ Handler = (*handler)(nil)

type Handler interface {
	i()
}

type handler struct {
}

func (h *handler) i() {}

//返回handler 对象
func New() Handler {
	return &handler{}
}
