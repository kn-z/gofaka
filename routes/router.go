package routes

import (
	"github.com/gin-gonic/gin"
	v1 "gofaka/api/v1"
	"gofaka/middleware"
	"gofaka/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	private := r.Group("api/v1")
	private.Use(middleware.JwtToken())
	{
		//user module routing interface
		private.POST("user/add", v1.AddUser)
		private.PUT("user/:id", v1.EditUser)
		private.DELETE("user/:id", v1.DeleteUser)

		//category module routing interfac
		private.POST("category/add", v1.AddCategory)
		private.PUT("category/:id", v1.EditCategory)
		private.DELETE("category/:id", v1.DeleteCategory)

		//article module routing interface
		private.POST("article/add", v1.AddArticle)
		private.PUT("article/:id", v1.EditArticle)
		private.DELETE("article/:id", v1.DeleteArticle)

	}

	public := r.Group("api/v1")
	{
		public.GET("users", v1.GetUsers)
		public.GET("category", v1.GetCategory)
		public.GET("article", v1.GetArticle)
		public.GET("article/list/:cid", v1.GetCateArt)
		public.GET("article/:id", v1.GetArtInfo)
	}

	r.Run(utils.HttpPort)
}
