package v1

import (
	"fmt"
	"github.com/go-gomail/gomail"
	"gofaka/utils"
	"gofaka/utils/errmsg"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type VerificationCode struct {
	code string
	time int64
}

var ch = make(chan *gomail.Message)
var UserMap = make(map[string]VerificationCode)

func SetMail(To []string) int {
	for _, to := range To {
		m := gomail.NewMessage()
		m.SetHeader("From", utils.WebName+"<"+utils.EmUser+">")
		m.SetHeader("To", to)
		m.SetHeader("Subject", utils.WebName+"邮箱验证码")
		timestamp := time.Now().Unix()
		UserMap[to] = VerificationCode{code: fmt.Sprintf("%06v", rand.New(rand.NewSource(timestamp)).Intn(1000000)), time: timestamp}
		m.SetBody("text/html", getBody(UserMap[to].code))
		ch <- m
	}
	return errmsg.SUCCESS
}

func SendMail() {
	d := gomail.NewDialer(utils.EmHost, utils.EmPort, utils.EmUser, utils.EmPasswd)
	var s gomail.SendCloser
	var err error
	open := false
	for {
		select {
		case m, ok := <-ch:
			if !ok {
				return
			}
			if !open {
				if s, err = d.Dial(); err != nil {
					panic(err)
				}
				open = true
			}
			if err := gomail.Send(s, m); err != nil {
				log.Println("Start")
				log.Print(err)
			}
		// Close the connection to the SMTP server if no email was sent in
		// the last 30 seconds.
		case <-time.After(30 * time.Second):
			if open {
				if err := s.Close(); err != nil {
					panic(err)
				}
				open = false
			}
		}
	}
}

func getBody(code string) string {
	content, err := os.ReadFile("../web/src/assets/static/verify.html")
	if err != nil {
		log.Println("Failed to read verify.html", err)
		panic(err)
	}

	body := string(content)
	body = strings.Replace(body, "{{$code}}", code, -1)
	body = strings.Replace(body, "{{$name}}", utils.WebName, -1)
	body = strings.Replace(body, "{{$url}}", utils.WebUrl, -1)
	body = strings.Replace(body, "{{$background-color}}", utils.WebBpColor, -1)
	return body
}

func CheckVerificationCode(email string, code string, second int64) int {
	if UserMap[email].code != code || len(code) < 6 {
		return errmsg.ErrorVerificationCodeError
	}
	if UserMap[email].time-time.Now().Unix() > second {
		delete(UserMap, email)
		return errmsg.ErrorVerificationCodeExpired
	}
	return errmsg.SUCCESS
}
