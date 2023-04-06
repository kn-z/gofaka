package v1

import (
	"github.com/gin-gonic/gin"
	"gofaka/middleware"
	"gofaka/model"
	"gofaka/utils/errmsg"
	"net/http"
	"strconv"
)

var code int

// add
func AddUser(c *gin.Context) {
	type Result struct {
		model.User
		Verify string `json:"verify"`
	}
	var data Result
	_ = c.ShouldBindJSON(&data)
	code = model.CheckEmail(data.Email)
	if code == errmsg.SUCCESS {
		code = CheckVerificationCode(data.Email, data.Verify, 300)
		if code == errmsg.SUCCESS {
			model.CreateUser(&data.User)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

// add
func AddUserByAdmin(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code = model.CheckEmail(data.Email)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

// query one
func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetUserByID(id)
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

// query list
func GetAllUser(c *gin.Context) {
	type Result struct {
		Users interface{} `json:"users"`
		Count int         `json:"count"`
	}
	var result Result
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	sortType := c.Query("sortType")
	sortKey := c.Query("sortKey")
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	result.Users, result.Count = model.GetUsers(pageSize, pageNum, sortType, sortKey)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

// edit
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	data.ID = uint(id)
	code = model.EditUser(id, &data)
	if code == errmsg.ERROR {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code)})
}

// delete
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// send mail
func SendEmail(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	to := []string{data.Email}
	code = SetMail(to)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetUserInfo(c *gin.Context) {
	var data model.User
	code = errmsg.SUCCESS
	data, code := middleware.GetUserByToken(c)
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
