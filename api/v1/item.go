package v1

import (
	"fmt"
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
	fmt.Println(data.Price)
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

//query list
func GetItems(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	data := model.GetItemList(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

//edit
func EditItem(c *gin.Context) {
	var data model.Item
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := model.EditItem(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code)})
	//
}

//delete
func DeleteItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteItem(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
