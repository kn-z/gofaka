package v1

import (
	"github.com/gin-gonic/gin"
	"gofaka/model"
	"gofaka/utils/errmsg"
	"net/http"
)

func AddApiKey(c *gin.Context) {
	var data model.Apikey
	_ = c.ShouldBindJSON(data)
	code = model.CreateApikey(&data)
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}
