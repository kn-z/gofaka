package v1

import (
	"github.com/gin-gonic/gin"
	"gofaka/model"
	"gofaka/utils/errmsg"
	"net/http"
	"strconv"
)

func AddPayment(c *gin.Context) {
	var data model.Payment
	_ = c.ShouldBindJSON(&data)
	code := model.CreatePayment(&data)
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

func GetPayment(c *gin.Context) {
	var data model.Payment
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetPayment(id)
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

func EditPayment(c *gin.Context) {
	var data model.Payment
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := model.EditPayment(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

func GetPaymentList(c *gin.Context) {
	var data []model.Payment
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	data, code := model.GetPaymentList(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

func GetAllPayment(c *gin.Context) {
	var data []model.Payment
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	data, code := model.GetAllPayment(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}
