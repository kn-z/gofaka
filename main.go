package main

import (
	"gofaka/api/v1"
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
	go v1.SendMail()
	//clean expired order
	v1.OrderClean()
	//init router
	routes.InitRouter()
}
