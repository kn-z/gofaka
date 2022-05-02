package routes

import (
	"github.com/gin-gonic/gin"
	v1 "gofaka/api/v1"
	"gofaka/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		//user module routing interface
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUserList)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)
		
		//category module routing interfac

		//article module routing interface
	}
	r.Run(utils.HttpPort)
}
