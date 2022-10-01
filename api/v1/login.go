package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gofaka/middleware"
	"gofaka/model"
	"gofaka/utils/errmsg"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	var token string

	_ = c.ShouldBindJSON(&data)

	code := model.CheckLogin(data.Email, data.Password)
	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.Email)
	}
	c.JSON(http.StatusOK, gin.H{
		"token":   token,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

func Pay(c *gin.Context) {
	qrUrl, code := middleware.PayTradePreCreate(c)
	c.JSON(http.StatusOK, gin.H{
		"qrUrl":   qrUrl,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

func Notify(c *gin.Context) {
	code := middleware.PayNotifyVerify(c)
	fmt.Println(code)
}
