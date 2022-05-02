package main

import (
	"gofaka/model"
	"gofaka/routes"
)

func main() {

	//database
	model.InitDb()

	routes.InitRouter()

}
