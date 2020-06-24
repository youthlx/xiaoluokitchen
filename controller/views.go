package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"xiaoluokitchen/common/exception"
	"xiaoluokitchen/common/resp"
	"xiaoluokitchen/model/vo"
)

func CheckLive(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
	})
}

func UserRegister(c *gin.Context) {
	var data *vo.UserInfo
	if err := c.ShouldBind(data); err != nil {
		resp.FailResp(c, exception.ParamCheckException)
		return
	}
	err := sg.userSrv.Register(context.Background(), data)
	if err != nil {
		resp.FailResp(c, exception.SystemException)
		return
	}
	resp.SuccessResp(c, data)
}
