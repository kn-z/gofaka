package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"gofaka/utils"
	"gofaka/utils/errmsg"
)

func PayTradePreCreate(c *gin.Context, OutTradeNo string, TotalAmount int) (string, int) {
	client, _ := alipay.NewClient(utils.AppId, utils.PrivateKey, true) //isP is sandbox
	bm := make(gopay.BodyMap)
	bm.Set("subject", "宝支付测")
	// out_trade_no is unique, an out_trade_no only be paid once
	bm.Set("out_trade_no", OutTradeNo)
	bm.Set("total_amount", float32(TotalAmount)/100)
	bm.Set("notify_url", utils.NotifyUrl)

	aliRsp, _ := client.TradePrecreate(c, bm)

	return aliRsp.Response.QrCode, errmsg.SUCCESS
}

func PayNotifyVerify(c *gin.Context) (gopay.BodyMap, int) {
	notifyReq, _ := alipay.ParseNotifyToBodyMap(c.Request)
	ok, _ := alipay.VerifySign(utils.AliPublicKey, notifyReq)
	if !ok {
		return notifyReq, errmsg.ERROR
	}
	return notifyReq, errmsg.SUCCESS
}
