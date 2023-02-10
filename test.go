package main

import (
	"fmt"
)

// mail template
//var ch = make(chan *gomail.Message)
//
//func SetMail(To []string) {
//	m := gomail.NewMessage()
//	m.SetHeader("From", "kitnoobcloud@gmail.co")
//	m.SetHeader("To", To...)
//	m.SetAddressHeader("Cc", "zw6979014@gmail.com", "zw")
//	m.SetHeader("Subject", "Hello!")
//	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
//	ch <- m
//	// prevent main process stop
//	time.Sleep(30 * time.Second)
//	//input := bufio.NewScanner(os.Stdin)
//	//for input.Scan() {
//	//	m := gomail.NewMessage()
//	//	m.SetHeader("From", "kitnoobcloud@gmail.co")
//	//	m.SetHeader("To", To...)
//	//	m.SetAddressHeader("Cc", "zw6979014@gmail.com", "zw")
//	//	m.SetHeader("Subject", "Hello!")
//	//	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
//	//	out <- m
//	//}
//}
//
//func SendMail() {
//	//m.Attach("/home/Alex/lolcat.jpg")
//	d := gomail.NewDialer(utils.EmHost, utils.EmPort, utils.EmUser, utils.EmPasswd)
//	var s gomail.SendCloser
//	var err error
//	open := false
//	for {
//		select {
//		case m, ok := <-ch:
//			if !ok {
//				return
//			}
//			if !open {
//				if s, err = d.Dial(); err != nil {
//					panic(err)
//				}
//				open = true
//			}
//			if err := gomail.Send(s, m); err != nil {
//				log.Println("Start")
//				log.Print(err)
//			}
//		// Close the connection to the SMTP server if no email was sent in
//		// the last 30 seconds.
//		case <-time.After(30 * time.Second):
//			if open {
//				if err := s.Close(); err != nil {
//					panic(err)
//				}
//				open = false
//			}
//		}
//	}
//}
//
//func main() {
//	utils.Init()
//	go SendMail()
//	SetMail([]string{"zw6979014@gmail.com"})
//
//	//SetMail(ch)
//	// Use the channel in your program to send emails.
//
//	// Close the channel to stop the mail daemon.
//	//close(ch)
//
//}

func main() {
	var list = []int{1, 2, 3, 4}
	for _, i := range list {
		i = 1
		fmt.Println(i)
		//fmt.Println(list[idx])
		//list[idx] = 1

	}
	fmt.Println(list)
}
