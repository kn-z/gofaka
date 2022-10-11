package middleware

import (
	"github.com/go-gomail/gomail"
	"gofaka/utils"
	"log"
	"os"
	"strings"
	"time"
)

var ch = make(chan *gomail.Message)

func SetMail(To []string) {
	m := gomail.NewMessage()
	m.SetHeader("From", utils.WebName+"<"+utils.EmUser+">")
	m.SetHeader("To", To...)
	m.SetHeader("Subject", utils.WebName+"邮箱验证码")
	m.SetBody("text/html", initBody())
	ch <- m
}

func SendMail() {
	//m.Attach("/home/Alex/lolcat.jpg")
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

func initBody() string {
	content, err := os.ReadFile("./utils/views/verify.html")
	if err != nil {
		log.Println("Failed to read verify.html", err)
		panic(err)
	}

	body := string(content)
	body = strings.Replace(body, "{{$code}}", "1234", -1)
	body = strings.Replace(body, "{{$name}}", utils.WebName, -1)
	body = strings.Replace(body, "{{$url}}", utils.WebUrl, -1)
	body = strings.Replace(body, "{{$background-color}}", utils.WebBpColor, -1)
	return body
}
