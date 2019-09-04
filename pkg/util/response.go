package util

import (
	"GinApi/middleware/error"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Context map[string]interface{}

//json返回格式
func JsonResponse(c *gin.Context, httpCode int, Context map[string]interface{}) {
	c.JSON(httpCode, Context)
}

//成功
func JsonSuccessResponse(c *gin.Context, errCode int, data interface{}) {

	Context = make(map[string]interface{})
	Context["code"] = errCode
	Context["msg"] = error.GetMsg(errCode)
	Context["data"] = data

	JsonResponse(c, http.StatusOK, Context)
}

//失败
func JsonErrResponse(c *gin.Context, errCode int) {
	Context = make(map[string]interface{})

	Context["code"] = errCode
	Context["msg"] = error.GetMsg(errCode)
	JsonResponse(c, http.StatusOK, Context)
}
