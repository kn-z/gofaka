package v1

import (
	"github.com/gin-gonic/gin"
	"gofaka/model"
	"gofaka/utils/errmsg"
	"net/http"
)

func GetOverride(c *gin.Context) {
	var data interface{}
	todayIncome, todayCount := model.GetTodayOrder()
	monthIncome, monthCount := model.GetMonthOrder()
	lastMonthIncome, lastMonthCount := model.GetLastMonthOrder()
	data = map[string]float64{
		"today_income":      todayIncome,
		"today_count":       float64(todayCount),
		"month_income":      monthIncome,
		"month_count":       float64(monthCount),
		"last_month_income": lastMonthIncome,
		"last_month_count":  float64(lastMonthCount),
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"code":    200,
		"message": errmsg.GetErrMsg(200),
	})
}
