package routes

import (
	"github.com/gin-gonic/gin"
	v1 "gofaka/api/v1"
	"gofaka/middleware"
	"gofaka/utils"
)

//首字母大写public
func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	private := r.Group("api/v1")
	private.Use(middleware.JwtToken())
	{
		//user module routing interface
		private.PUT("user/:id", v1.EditUser)
		private.GET("users", v1.GetUsers)
		private.DELETE("user/:id", v1.DeleteUser)

		//category module routing interfac
		private.POST("category/add", v1.AddCategory)
		private.PUT("category/:id", v1.EditCategory)
		private.DELETE("category/:id", v1.DeleteCategory)

		//article module routing interface
		private.POST("article/add", v1.AddArticle)
		private.PUT("article/:id", v1.EditArticle)
		private.DELETE("article/:id", v1.DeleteArticle)

		//item module routing interface
		private.POST("item/add", v1.AddItem)
		private.PUT("item/:id", v1.EditItem)
		private.DELETE("item/:id", v1.DeleteItem)

		private.POST("upload", v1.UpLoad)
		//private.POST("upload", func(c *gin.Context) {
		//	file, err := c.FormFile("file")
		//	if err != nil {
		//		c.String(500, "上传文件出错")
		//	}
		//	c.SaveUploadedFile(file, file.Filename)
		//	c.String(http.StatusOK, file.Filename+"上传成功")
		//})

	}

	public := r.Group("api/v1")
	{
		public.POST("notify", v1.Notify)
		public.POST("pay", v1.Pay)

		//user module routing interface
		public.POST("login", v1.Login)
		public.POST("user/add", v1.AddUser)

		//item module routing interface
		public.GET("items", v1.GetItems)

		public.GET("category", v1.GetCategory)
		public.GET("article", v1.GetArticle)
		public.GET("article/list/:cid", v1.GetCateArt)
		public.GET("article/:id", v1.GetArtInfo)

	}

	r.Run(utils.HttpPort)
}
