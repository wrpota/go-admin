package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/wrpota/go-gin/internal/global/variable"
	"github.com/wrpota/go-gin/internal/model"
	"github.com/wrpota/go-gin/internal/pkg/response"
	"golang.org/x/crypto/bcrypt"
)

type LoginResuqet struct {
	UserName string `json:"username" form:"username" query:"username"`
	Password string `json:"password" form:"password" query:"password"`
}

type loginResponse map[string]interface{}

func (u *handler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginForm LoginResuqet
		c.Bind(&loginForm)
		var admin model.Admin
		if result := variable.GormWriteDb.Where("username=?", loginForm.UserName).First(&admin); result.Error != nil {
			response.Fail(c, 403, "用户名不存在", "")
			return
		}
		// result, _ := bcrypt.GenerateFromPassword([]byte(loginForm.Password), bcrypt.DefaultCost)
		if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(loginForm.Password)); err != nil {
			response.Fail(c, 403, "密码错误", "")
			return
		}
		//颁发token
		token, err := CreateToken(&admin)
		if err != nil {
			response.Fail(c, 403, "登录失败", "")
		}
		//返回数据
		user := loginResponse{}
		user["token"] = token
		user["id"] = admin.ID
		user["username"] = admin.Username

		response.Success(c, "登录成功", user)
	}
}
