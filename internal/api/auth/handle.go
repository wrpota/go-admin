package auth

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/wrpota/go-gin/configs"
	"github.com/wrpota/go-gin/internal/middleware"
	"github.com/wrpota/go-gin/internal/model"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()
	Login() gin.HandlerFunc
	RefreshToken() gin.HandlerFunc
}

type handler struct {
}

func (h *handler) i() {}

//返回handler 对象
func New() Handler {
	return &handler{}
}

func CreateToken(admin *model.Admin) (string, error) {
	claims := middleware.AuthClaims{
		Name: admin.Username,
		ID:   admin.ID,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix()),                       // 签名生效时间
			ExpiresAt: int64(time.Now().Add(time.Minute * 30).Unix()), // 签名过期时间
			Issuer:    configs.Get().GetString("App.Name"),            // 签名颁发者
		},
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(configs.Get().GetString("App.AccessSecret")))
	if err != nil {
		return "", err
	}
	return token, nil
}
