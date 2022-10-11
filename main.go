package main

import (
	"gofaka/middleware"
	"gofaka/model"
	"gofaka/routes"
	"gofaka/utils"
)

func main() {
	//init config
	utils.Init()
	//init database
	model.InitDb()
	//daemon of mail
	go middleware.SendMail()
	//init router
	routes.InitRouter()
}
