package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"gofaka/utils/errmsg"
	"time"
)

var (
	appId = "2021002161609531"
	//支付宝公钥，不是生成的
	aliPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA3cFmJ7tfJs1l673JxvHwAuuzqy8LgFj4G9Zo5CFGPVs2XRs4cSmILDDm1Cp8QsQv3Tkad9TSwCB5DrAzUEkRteuqTRWU+7LvrGDzraJHM3glkgbI9NmiQe3LlQEayclfd8gvJEMPE2khCB48R4gCnT1QzS47j5Q4NO+lyho8W54rsnuYIqNBpL8p1wSXsYdVDumY3ZpErQ64sS7b0rDMK9xFMyLcON5wiSS+p1yE4G/QBRhTrMRuW5JFAXpaoaOwR0sVqZH/UubPPYHtWn8FoAsTlUWt1cP2D3CR2PdUVrDbmE/P/BQBjVvFI+PIqyhMt72gQDGKYJYwQ2Va5JB0RwIDAQAB"
	privateKey   = "MIIEogIBAAKCAQEAhAJZEhiQy/ASCa8ZLWA1x6CsR/eRUpm/Px/toEwVC6jnJg5VbpI1/NM6NATODmjClY29F9aXXFlyirVBHfvhGav5IPGXEyIisojHRxAs1Jucz0PBRFjEh8c2mTenJYcnmvaQ7bTsV8CM7XkyzpbcMBcnC76dzDlKrJMgwj1rI6eGVz2p7AsqHo4HoIVYydWeGjGjC7TVaZDgxP6JUdVyMFNrgrZj+/wvKDiXbtpvlGJJTEf2zpZj0PUW5keRbnFRRupmY3FE3lmikgiGB2iOJNwAGsstXvslFUw4YEUuODb41PIjyGmRjInVeOH881SVDLy+pgF7GDk/gR3kYFB+nwIDAQABAoIBAAsNY0W3ls/sTqZO6a2542bOVf5EhP9EbhWr56pHFHAAeTHfe1mhljGPwcy/Jj1gmgIJdu57AfsyZfulB0mqYANVOat/bqWkcwE9oGmbuhUm3i8gPhspz3KKxFB5r13d/fvkbufoAO2r6mCriAxx5weDuLosAGwr08u1GcYJfGHksvRleFtMDYKIgQe/dvRfBf0ObRzNwHnMW/pciaJmtQovi2cfZbmOxnc73f++oNUwfLl7kqVt9QAX3ddOZkrQHudjwH3oOGqkSC7kIiqa1QxZogZ+476EW5M8oV5h5fsmAPISj9a6M0LqQknZwbo2xF6aKjQFry5wSbMSHOZHOUECgYEAv7T7DF7dAg/ZcDy5BjkiW2GyLoy87piIbE9pt8QIg7fAnGC3ApoXnu2ioMYGIUXSlklPHFAB1DYMTULTpSSCmTwtaF5o9LwC2QG2WYzFqWqoQVQsCTBbKfa9RiJEjenI5SyhEngeXOs8iT5YpnQvF1BNQI+aiJfPehwnnVi3/T8CgYEAsEgCGiIWF1PNSujKYHnu6qGGHnO7DKzh7X9O7bCu3Vp0CRhBf6n/zyrBBBFbBUyWPC5NJ4XgiwfPXQQQzF3J38TRPxLbIprFgnPMsGMTQbQm2ciVzIsvZXd5gMZdXF+ZmimTdn8WcEAv5Sk5n3e0gqoTK3+F0ujyX2NibUTlRqECgYAuEAteXpDWBP7nBAAlKadCs8e/fZuL7OSiubYaLKUrGQTTzj1LB8FzM4AnB03DwuYlrDmxANxfpBjym4MFJC+pKBd1A3JOk7pPcCTjgXqhCXqiL9pg3tiYzauO5X75ZloaDs4pBOmuw+sIww1D+ZizDl1xjM/B0FBO8+Lk4MNcuwKBgC8MySLYfjTztRONVpaxdMdDHVz7Xq1fZ13QYOyn/8Qs5FOZGcJNSW2t556CU1zyuBaP9R/bZ7cz+nDFKQai8cK78W14RuzRim3rInLhvr4Gq2ftVa4maBwY62EnkLua+JBhEG7MNNz5BM+RVUPu20sUwdEWVE2axzYWfKrVfKyBAoGAc+kCLZgI6DMrKrBGjISUyo3nTlnpP0yEwtAIUWBMkZwJIKUiXaeOuigTLaiiiOTtxOtumznj/OMs+y4nUE+hRY3zLQ0tG0TcnZzD4V14Ge64a9aUigf14afzw8q6y4xTA1jQr356+VDJyBMpidu0ee8bc0jhoXD+uuXc1AgoLjA="
	client, _    = alipay.NewClient(appId, privateKey, true) //isP is sandbox
)

func PayTradePreCreate(c *gin.Context) (string, int) {
	bm := make(gopay.BodyMap)
	bm.Set("subject", "宝支付测")
	// out_trade_no is unique, an out_trade_no only be paid once

	outTradeNo := time.Now().Format("20060102150405.999")
	outTradeNo = outTradeNo[:14] + outTradeNo[15:]
	bm.Set("out_trade_no", outTradeNo)
	bm.Set("total_amount", "0.01")
	bm.Set("notify_url", "http://hk.knyun.xyz:3000/api/v1/notify")

	aliRsp, _ := client.TradePrecreate(c, bm)

	return aliRsp.Response.QrCode, errmsg.SUCCESS
}

func PayNotifyVerify(c *gin.Context) int {
	notifyReq, _ := alipay.ParseNotifyToBodyMap(c.Request)
	fmt.Println(notifyReq)
	ok, _ := alipay.VerifySign(aliPublicKey, notifyReq)
	if !ok {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
