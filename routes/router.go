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
		router.GET("users", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)

		//category module routing interfac
		router.POST("category/add", v1.AddCategory)
		router.GET("category", v1.GetCategory)
		router.PUT("category/:id", v1.EditCategory)
		router.DELETE("category/:id", v1.DeleteCategory)

		//article module routing interface
		router.POST("article/add", v1.AddArticle)
		router.GET("article", v1.GetArticle)
		router.GET("article/list/:cid", v1.GetCateArt)
		router.GET("article/:id", v1.GetArtInfo)
		router.PUT("article/:id", v1.EditArticle)
		router.DELETE("article/:id", v1.DeleteArticle)

	}

	r.Run(utils.HttpPort)
}
