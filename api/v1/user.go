package v1

import (
	"github.com/gin-gonic/gin"
	"gofaka/model"
	"gofaka/utils/errmsg"
	"net/http"
)

var code int

//query exist
func UserExist(c *gin.Context) {
	//
}

//add
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

//query one

//query list
func GetUserList(c *gin.Context) {
	//
}

//edit
func EditUser(c *gin.Context) {
	//
}

//delete
func DeleteUser(c *gin.Context) {
	//
}
