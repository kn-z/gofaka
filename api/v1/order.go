package v1

import (
	"github.com/gin-gonic/gin"
	"gofaka/middleware"
	"gofaka/model"
	"gofaka/utils/errmsg"
	"net/http"
	"strconv"
	"time"
)

func CreateOrder(c *gin.Context) {
	var data model.Order
	_ = c.ShouldBindJSON(&data)
	code = model.CreateOrder(&data)
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
	go OrderExpired(&data)
}

func OrderExpired(order *model.Order) {
	select {
	case <-time.After(time.Second * 120):
		_ = model.CancelOrder(order)
	}
}

func GetOrderInfo(c *gin.Context) {
	outTradeNo := c.Param("outTradeNo")
	data, code := model.GetOrderInfo(outTradeNo)
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

// query list
func GetAllOrder(c *gin.Context) {
	type Result struct {
		Orders interface{} `json:"orders"`
		Count  int         `json:"count"`
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

	result.Orders, result.Count = model.GetAllOrder(pageSize, pageNum, sortType, sortKey)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

// edit
func EditOrder(c *gin.Context) {
	var data model.Order
	//id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := model.EditOrder(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code)})
	//
}

// cancel
func CancelOrder(c *gin.Context) {
	var data model.Order
	_ = c.ShouldBindJSON(&data)
	code := model.CancelOrder(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code)})
	//
}

// pay
func PayOrder(c *gin.Context) {
	var data model.Order
	_ = c.ShouldBindJSON(&data)
	data, code := model.GetOrder(data.OutTradeNo)
	if data.Status != 0 {
		code = errmsg.ErrorOrderInvalid
	}
	qrUrl := ""
	if code == errmsg.SUCCESS {
		qrUrl, code = PayTradePreCreate(c, data.OutTradeNo, data.TotalAmount)
	}
	c.JSON(http.StatusOK, gin.H{
		"qrUrl":   qrUrl,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

// check
func CallBackOrder(c *gin.Context) {

	notifyReq, code := PayNotifyVerify(c)
	if code == errmsg.SUCCESS {
		code = model.CallBackOrder(notifyReq["out_trade_no"].(string))
	}
	c.JSON(http.StatusOK, gin.H{
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

func GetUserOrder(c *gin.Context) {
	var user model.User
	var data interface{}
	user, code = middleware.GetUserByToken(c)
	status, _ := strconv.Atoi(c.Query("status"))
	if code == 200 {
		data = model.GetUserOrderByToken(user.Email, status)
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

func QueryOrder(c *gin.Context) {
	var data interface{}
	email := c.Query("email")
	status, _ := strconv.Atoi(c.Query("status"))
	if len(email) == 0 {
		code = errmsg.ErrorInvalidEmail
	} else {
		data = model.GetOrderByEmail(email, status)
		code = errmsg.SUCCESS
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

func OrderClean() {
	model.CancelExpiredOrder()
}
