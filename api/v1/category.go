package v1

import (
	"github.com/gin-gonic/gin"
	"gofaka/model"
	"gofaka/utils/errmsg"
	"net/http"
	"strconv"
)

// add
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		model.CreateCategory(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

// query list
func GetCategoryInfo(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetCategoryInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

// query list
func GetCategoryList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	//if pageSize == 0 {
	//	pageSize = -1
	//}
	//if pageNum == 0 {
	//	pageNum = -1
	//}
	data := model.GetCategory(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

// query list
func GetAllCategory(c *gin.Context) {
	data := model.GetAllCategory()
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}

// edit
func EditCategory(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		model.EditCategory(id, &data)
	}
	if code == errmsg.ERROR {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
	//
}

// delete
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
