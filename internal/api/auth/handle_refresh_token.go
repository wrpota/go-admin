package auth

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wrpota/go-gin/internal/global/variable"
	"github.com/wrpota/go-gin/internal/model"
	"github.com/wrpota/go-gin/internal/pkg/response"
)

func Strval(value interface{}) string {
	var str string
	if value == nil {
		return str
	}
	// vt := value.(type)
	switch value.(type) {
	case float64:
		ft := value.(float64)
		str = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		str = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		str = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		str = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		str = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		str = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		str = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		str = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		str = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		str = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		str = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		str = strconv.FormatUint(it, 10)
	case string:
		str = value.(string)
	case []byte:
		str = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		str = string(newValue)
	}

	return str
}
func (u *handler) RefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authId, hasClaims := c.Get("authId")
		if !hasClaims {
			response.Fail(c, http.StatusForbidden, "请重新登录", "")
			return
		}
		id, ok := authId.(int64)
		if !ok {
			response.Fail(c, http.StatusForbidden, "请重新登录", "")
			return
		}

		var admin model.Admin
		if result := variable.GormReadDb.Where("id = ?", id).First(&admin); result.Error != nil {
			response.Fail(c, 403, "用户名不存在", "")
			return
		}
		//颁发token
		token, err := CreateToken(&admin)
		if err != nil {
			response.Fail(c, http.StatusForbidden, "刷新失败", "")
		}
		//返回数据
		user := loginResponse{}
		user["token"] = token
		user["id"] = admin.ID
		user["username"] = admin.Username

		response.Success(c, "刷新成功", user)
	}
}
