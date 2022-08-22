package user

import (
	"github.com/gin-gonic/gin"
	"github.com/wrpota/go-gin/internal/pkg/response"
)

type UserRequest struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

func (u *handler) HelloWorld() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := new(UserRequest)
		if err := c.Bind(user); err != nil {
			return
		}
		response.Success(c, "success", user)
	}
}
