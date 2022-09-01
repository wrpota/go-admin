package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/wrpota/go-gin/configs"
	"github.com/wrpota/go-gin/internal/pkg/response"
)

var (
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("Couldn't handle this token:")
)

type AuthClaims struct {
	Name string `json:"name,omitempty"`
	ID   int64  `json:"id,omitempty"`
	jwt.StandardClaims
}

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(configs.Get().GetString("App.AccessSecret")),
	}
}

func (j *JWT) ParserToken(tokenString string) (*AuthClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenMalformed
			}

		}
	}

	// 将token中的claims信息解析出来并断言成用户自定义的有效载荷结构
	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, TokenMalformed

}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 通过http header中的token解析来认证
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			response.Fail(c, http.StatusForbidden, "未登录", "")
			return
		}
		headerList := strings.Split(authorization, " ")
		if len(headerList) != 2 {
			response.Fail(c, http.StatusBadRequest, "无法解析 Authorization 字段", "")
			return
		}
		t := headerList[0]
		token := headerList[1]
		if t != "Bearer" {
			response.Fail(c, http.StatusBadRequest, "认证类型错误, 当前只支持 Bearer", "")
			return
		}
		// 初始化一个JWT对象实例，并根据结构体方法来解析token
		jwt := NewJWT()
		// 解析token中包含的相关信息(有效载荷)
		claims, err := jwt.ParserToken(token)

		if err != nil {
			// token过期
			if err == TokenExpired {
				response.Fail(c, http.StatusForbidden, "token授权已过期，请重新申请授权", "")
				return
			}
			// 其他错误
			response.Fail(c, http.StatusForbidden, err.Error(), "")
			return
		}

		// 将解析后的有效载荷claims重新写入gin.Context引用对象中
		c.Set("authId", claims.ID)

		c.Next()
	}
}

func (j *JWT) CreateToken(claims AuthClaims) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return at.SignedString([]byte(configs.Get().GetString("App.AccessSecret")))
}
