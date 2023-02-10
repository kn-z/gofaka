package v1

import (
	"github.com/gin-gonic/gin"
	"gofaka/model"
	"gofaka/utils/errmsg"
	"net/http"
	"strconv"
)

func AddItem(c *gin.Context) {
	var data model.Item
	_ = c.ShouldBindJSON(&data)
	code = model.CreateItem(&data)
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

func GetItem(c *gin.Context) {
	var data model.Item
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetItem(id)
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

// query list
func GetAllItem(c *gin.Context) {
	type Result struct {
		Items interface{} `json:"items"`
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
	result.Items, result.Count = model.GetAllItem(pageSize, pageNum, sortType, sortKey)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

func GetItemByOrder(c *gin.Context) {
	outTradeNo := c.Param("outTradeNo")
	data := model.GetItemListByOrder(outTradeNo)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

// edit
func EditItem(c *gin.Context) {
	var data model.Item
	_ = c.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.EditItem(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code)})
}

// delete
func DeleteItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteItem(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func LockItems2Order(c *gin.Context) {
	var data model.Order
	_ = c.ShouldBindJSON(&data)
	code := model.LockItems2Order(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code)})
}

func UnlockItems2Order(c *gin.Context) {
	var data model.Order
	_ = c.ShouldBindJSON(&data)
	code := model.UnlockItems2Order(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code)})
}
