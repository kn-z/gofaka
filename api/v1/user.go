package v1

import (
	"github.com/gin-gonic/gin"
	"gofaka/model"
	"gofaka/utils/errmsg"
	"net/http"
	"strconv"
)

var code int

//add
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Email)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
		//fmt.Println(data.Email)
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

//query one

//query list
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	//if pageSize == 0 {
	//	pageSize = -1
	//}
	//if pageNum == 0 {
	//	pageNum = -1
	//}
	data := model.GetUsers(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

//edit
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Email)
	if code == errmsg.SUCCESS {
		model.EditUser(id, &data)
	}
	if code == errmsg.ERROR {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code)})
	//
}

//delete
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//send mail
func SendEmail(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code := model.SendEmail(data.Email)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
