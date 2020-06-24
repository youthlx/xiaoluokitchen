package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xiaoluokitchen/common/exception"
)

type JsonResp struct {
	code    int         `json:"code"`
	message string      `json:"message"`
	data    interface{} `json:"data,omitempty"`
}

func FailResp(c *gin.Context, ex *exception.Exception) *JsonResp {
	jp := &JsonResp{
		code:    ex.ErrCode,
		message: ex.ErrMsg,
	}
	c.AbortWithStatusJSON(http.StatusOK, jp)
}

func SuccessResp(c *gin.Context, data interface{}) *JsonResp {
	jp := &JsonResp{
		code:    0,
		message: "SUCCESS",
		data:    data,
	}
	c.JSON(http.StatusOK, jp)
}

func (jp *JsonResp) Code() int {
	return jp.code
}

func (jp *JsonResp) Message() string {
	return jp.message
}

func (jp *JsonResp) Data() interface{} {
	return jp.message
}
