package v1

import (
	"github.com/gin-gonic/gin"
	"gofaka/utils/errmsg"
	"net/http"
)

func UpLoad(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		code = errmsg.ERROR
	}
	dst := "./upload/" + file.Filename
	_ = c.SaveUploadedFile(file, dst)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"data":    file.Filename,
		"message": errmsg.GetErrMsg(code),
		"status":  code,
	})
}
