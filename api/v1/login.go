package v1

import (
	"github.com/gin-gonic/gin"
	"gofaka/middleware"
	"gofaka/model"
	"gofaka/utils/errmsg"
	"net/http"
)

func Login(c *gin.Context) (int, string, int) {
	var data model.User
	var token string

	_ = c.ShouldBindJSON(&data)

	role, code := model.CheckLogin(data.Email, data.Password)
	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.Email)
	}
	c.Abort()
	return role, token, code
}

func UserLogin(c *gin.Context) {
	_, token, code := Login(c)
	c.JSON(http.StatusOK, gin.H{
		"token":   token,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

func AdminLogin(c *gin.Context) {
	role, token, code := Login(c)
	if role != 1 {
		code = errmsg.ErrorUserNoRight
		token = ""
	}
	c.JSON(http.StatusOK, gin.H{
		"token":   token,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

//func Notify(c *gin.Context) {
//	code := middleware.PayNotifyVerify(c)
//	fmt.Println(code)
//}
