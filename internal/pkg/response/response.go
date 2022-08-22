package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReturnJson(Context *gin.Context, httpCode int, dataCode int, msg string, data interface{}) {
	Context.JSON(httpCode, gin.H{
		"code": dataCode,
		"msg":  msg,
		"data": data,
	})
}

//ReturnJsonFromString 将json字符串以标准json格式返回
func ReturnJsonFromString(Context *gin.Context, httpCode int, jsonStr string) {
	Context.Header("Content-Type", "application/json; charset=utf-8")
	Context.String(httpCode, jsonStr)
}

//Success 直接返回成功
func Success(c *gin.Context, msg string, data interface{}) {
	ReturnJson(c, http.StatusOK, http.StatusOK, msg, data)
}

//Fail 失败的业务逻辑
func Fail(c *gin.Context, dataCode int, msg string, data interface{}) {
	ReturnJson(c, http.StatusBadRequest, dataCode, msg, data)
	c.Abort()
}

func JsonNoContent(c *gin.Context, httpCode int) {
	ReturnJson(c, httpCode, http.StatusNoContent, "", nil)
}

// // ErrorSystem 系统执行代码错误
// func ErrorSystem(c *gin.Context, msg string, data interface{}) {
// 	ReturnJson(c, http.StatusInternalServerError, consts.ServerOccurredErrorCode, consts.ServerOccurredErrorMsg+msg, data)
// 	c.Abort()
// }
