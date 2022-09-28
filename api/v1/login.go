package v1

import (
	"github.com/gin-gonic/gin"
	"gofaka/middleware"
	"gofaka/model"
	"gofaka/utils/errmsg"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	var token string

	c.ShouldBindJSON(&data)

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
