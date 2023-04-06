package v1

import (
	"github.com/gin-gonic/gin"
	"gofaka/model"
	"gofaka/utils/errmsg"
	"net/http"
	"strconv"
)

func AddGoods(c *gin.Context) {
	var data model.Goods
	_ = c.ShouldBindJSON(&data)
	code := model.CreateGoods(&data)
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

// query list
func GetGoodsList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	data := model.GetGoodsList(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

// query list
func GetAllGoods(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	data := model.GetAllGoods(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

func GetGoods(c *gin.Context) {
	var data model.Goods
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetGoods(id)
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

// edit
func EditGoods(c *gin.Context) {
	var data model.Goods
	_ = c.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.EditGoods(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code)})
}

// delete
func DeleteGoods(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteGoods(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func CheckGoodsStock(c *gin.Context) {
	code = model.CheckGoodsStock()
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func SortGoods(c *gin.Context) {
	sort, _ := strconv.Atoi(c.Param("sort"))
	for data, _ := range sort {
		code = model.UpdateSort(data.id, data.sort)
		if code == errmsg.ERROR {
			break
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
