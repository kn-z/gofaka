package main

import (
	"fmt"
	"github.com/smartwalle/alipay"
	"net/http"
	"time"
)

var (
	appId = "2021002161609531"
	//支付宝公钥，不是生成的
	aliPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA3cFmJ7tfJs1l673JxvHwAuuzqy8LgFj4G9Zo5CFGPVs2XRs4cSmILDDm1Cp8QsQv3Tkad9TSwCB5DrAzUEkRteuqTRWU+7LvrGDzraJHM3glkgbI9NmiQe3LlQEayclfd8gvJEMPE2khCB48R4gCnT1QzS47j5Q4NO+lyho8W54rsnuYIqNBpL8p1wSXsYdVDumY3ZpErQ64sS7b0rDMK9xFMyLcON5wiSS+p1yE4G/QBRhTrMRuW5JFAXpaoaOwR0sVqZH/UubPPYHtWn8FoAsTlUWt1cP2D3CR2PdUVrDbmE/P/BQBjVvFI+PIqyhMt72gQDGKYJYwQ2Va5JB0RwIDAQAB"
	privateKey   = "MIIEogIBAAKCAQEAhAJZEhiQy/ASCa8ZLWA1x6CsR/eRUpm/Px/toEwVC6jnJg5VbpI1/NM6NATODmjClY29F9aXXFlyirVBHfvhGav5IPGXEyIisojHRxAs1Jucz0PBRFjEh8c2mTenJYcnmvaQ7bTsV8CM7XkyzpbcMBcnC76dzDlKrJMgwj1rI6eGVz2p7AsqHo4HoIVYydWeGjGjC7TVaZDgxP6JUdVyMFNrgrZj+/wvKDiXbtpvlGJJTEf2zpZj0PUW5keRbnFRRupmY3FE3lmikgiGB2iOJNwAGsstXvslFUw4YEUuODb41PIjyGmRjInVeOH881SVDLy+pgF7GDk/gR3kYFB+nwIDAQABAoIBAAsNY0W3ls/sTqZO6a2542bOVf5EhP9EbhWr56pHFHAAeTHfe1mhljGPwcy/Jj1gmgIJdu57AfsyZfulB0mqYANVOat/bqWkcwE9oGmbuhUm3i8gPhspz3KKxFB5r13d/fvkbufoAO2r6mCriAxx5weDuLosAGwr08u1GcYJfGHksvRleFtMDYKIgQe/dvRfBf0ObRzNwHnMW/pciaJmtQovi2cfZbmOxnc73f++oNUwfLl7kqVt9QAX3ddOZkrQHudjwH3oOGqkSC7kIiqa1QxZogZ+476EW5M8oV5h5fsmAPISj9a6M0LqQknZwbo2xF6aKjQFry5wSbMSHOZHOUECgYEAv7T7DF7dAg/ZcDy5BjkiW2GyLoy87piIbE9pt8QIg7fAnGC3ApoXnu2ioMYGIUXSlklPHFAB1DYMTULTpSSCmTwtaF5o9LwC2QG2WYzFqWqoQVQsCTBbKfa9RiJEjenI5SyhEngeXOs8iT5YpnQvF1BNQI+aiJfPehwnnVi3/T8CgYEAsEgCGiIWF1PNSujKYHnu6qGGHnO7DKzh7X9O7bCu3Vp0CRhBf6n/zyrBBBFbBUyWPC5NJ4XgiwfPXQQQzF3J38TRPxLbIprFgnPMsGMTQbQm2ciVzIsvZXd5gMZdXF+ZmimTdn8WcEAv5Sk5n3e0gqoTK3+F0ujyX2NibUTlRqECgYAuEAteXpDWBP7nBAAlKadCs8e/fZuL7OSiubYaLKUrGQTTzj1LB8FzM4AnB03DwuYlrDmxANxfpBjym4MFJC+pKBd1A3JOk7pPcCTjgXqhCXqiL9pg3tiYzauO5X75ZloaDs4pBOmuw+sIww1D+ZizDl1xjM/B0FBO8+Lk4MNcuwKBgC8MySLYfjTztRONVpaxdMdDHVz7Xq1fZ13QYOyn/8Qs5FOZGcJNSW2t556CU1zyuBaP9R/bZ7cz+nDFKQai8cK78W14RuzRim3rInLhvr4Gq2ftVa4maBwY62EnkLua+JBhEG7MNNz5BM+RVUPu20sUwdEWVE2axzYWfKrVfKyBAoGAc+kCLZgI6DMrKrBGjISUyo3nTlnpP0yEwtAIUWBMkZwJIKUiXaeOuigTLaiiiOTtxOtumznj/OMs+y4nUE+hRY3zLQ0tG0TcnZzD4V14Ge64a9aUigf14afzw8q6y4xTA1jQr356+VDJyBMpidu0ee8bc0jhoXD+uuXc1AgoLjA="
	client       = alipay.New(appId, aliPublicKey, privateKey, true) //isP is sanbox
)

func PayTradePreCreate() {
	pay := alipay.AliPayTradePreCreate{}
	// 支付成功之后，支付宝将会重定向到该 URL
	pay.ReturnURL = "http://www.baidu.com"
	pay.NotifyURL = "http://hk.knyun.xyz:8088/return"
	//支付标题
	pay.Subject = "支付宝支付测试"
	//订单号，一个订单号只能支付一次
	pay.OutTradeNo = time.Now().String()
	//销售产品码，与支付宝签约的产品码名称,目前仅支持FAST_INSTANT_TRADE_PAY
	//pay.ProductCode = "FAST_INSTANT_TRADE_PAY"
	//金额
	pay.TotalAmount = "0.01"
	response, err := client.TradePreCreate(pay)
	if err != nil {
		fmt.Println(err)
	}
	//payURL := url.String()
	//这个 payURL 即是用于支付的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	fmt.Println(response)

	//打开默认浏览器
	//payURL = strings.Replace(payURL, "&", "^&", -1)
	//exec.Command("cmd", "/c", "start", payURL).Start()
}

//手机客户端支付
func WapAlipay() {
	pay := alipay.AliPayTradeWapPay{}
	// 支付成功之后，支付宝将会重定向到该 URL
	pay.ReturnURL = "http://www.baidu.com"
	pay.NotifyURL = "http://hk.knyun.xyz:3000/api/v1/pay"
	//支付标题
	pay.Subject = "支付宝支付测试"
	//订单号，一个订单号只能支付一次
	pay.OutTradeNo = time.Now().String()
	//商品code
	//pay.ProductCode = time.Now().String()
	//金额
	pay.TotalAmount = "0.01"
	fmt.Println(pay.OutTradeNo)
	url, err := client.TradeWapPay(pay)
	if err != nil {
		fmt.Println(err)
	}
	payURL := url.String()
	//这个 payURL 即是用于支付的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	fmt.Println(payURL)
	//打开默认浏览器
	//payURL = strings.Replace(payURL, "&", "^&", -1)
	//exec.Command("cmd", "/c", "start", payURL).Start()
}

func main() {
	//生成支付URL
	//WapAlipay()
	PayTradePreCreate()
	//// 支付成功之后的返回URL页面
	http.HandleFunc("/return", func(rep http.ResponseWriter, req *http.Request) {
		rep.Write([]byte("等待"))
		req.ParseForm()
		ok, err := client.VerifySign(req.Form)
		if err == nil && ok {
			rep.Write([]byte("支付成功"))
		}
	})
	fmt.Println("server start....")
	http.ListenAndServe(":8088", nil)
}
