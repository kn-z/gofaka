package routes

import (
	"github.com/gin-gonic/gin"
	v1 "gofaka/api/v1"
	"gofaka/middleware"
	"gofaka/utils"
)

// 首字母大写public
func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	admin := r.Group("api/v1")
	admin.Use(middleware.AdminJwtToken())
	{
		//user module routing interface
		admin.GET("user/all", v1.GetAllUser)
		admin.PUT("user/:id", v1.EditUser)
		admin.GET("user/:id", v1.GetUser)
		admin.DELETE("user/:id", v1.DeleteUser)
		admin.POST("user/add", v1.AddUserByAdmin)

		//category module routing interface
		admin.GET("category/all", v1.GetAllCategory)
		admin.POST("category/add", v1.AddCategory)
		admin.PUT("category/:id", v1.EditCategory)
		admin.GET("category/:id", v1.GetCategoryInfo)
		admin.DELETE("category/:id", v1.DeleteCategory)

		//goods module routing interface
		admin.GET("goods/all", v1.GetAllGoods)
		admin.POST("good/add", v1.AddGoods)
		admin.PUT("good/:id", v1.EditGoods)
		admin.DELETE("good/:id", v1.DeleteGoods)
		admin.PUT("good/check", v1.CheckGoodsStock)
		admin.PUT("goods/sort", v1.SortGoods)

		//order module routing interface
		admin.GET("order/all", v1.GetAllOrder)
		admin.GET("orders", v1.GetUserOrder)

		//item module routing interface
		admin.GET("item/all", v1.GetAllItem)
		admin.PUT("item/:id", v1.EditItem)
		admin.GET("item/:id", v1.GetItem)
		admin.POST("item/add", v1.AddItem)
		admin.POST("item/add/batch", v1.BatchAddItem)

		//notice module routing interface
		//admin.GET("notice/all", v1.GetItemList)
		admin.GET("notice/all", v1.GetAllNotice)
		admin.PUT("notice/:id", v1.EditNotice)
		admin.POST("notice/add", v1.AddNotice)
		admin.DELETE("notice/:id", v1.DeleteNotice)

		admin.POST("upload", v1.UpLoad)
		//user.POST("upload", func(c *gin.Context) {
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
		public.POST("notify", v1.CallBackOrder)
		public.POST("verify", v1.SendEmail)

		//user module routing interface
		public.POST("login", v1.UserLogin)
		public.POST("admin", v1.AdminLogin)
		public.POST("user/create", v1.AddUser)
		public.POST("user/info", v1.GetUserInfo)

		//good module routing interface
		public.GET("goods/list", v1.GetGoodsList)
		public.GET("good/:id", v1.GetGoods)

		//order module routing interface
		public.GET("orders/search", v1.QueryOrder)
		public.GET("order/:outTradeNo", v1.GetOrderInfo)
		public.POST("order/create", v1.CreateOrder)
		public.PUT("order/cancel", v1.CancelOrder)

		//item module routing interface
		public.PUT("binditems", v1.LockItems2Order)
		public.PUT("unbinditems", v1.UnlockItems2Order)
		public.GET("items/:outTradeNo", v1.GetItemByOrder)

		//notice module routing interface
		public.GET("notice/list", v1.GetNoticeList)
		public.GET("notice/:id", v1.GetNoticeByID)

		//category module routing interface
		public.GET("category", v1.GetCategory)

		public.POST("pay", v1.PayOrder)

	}

	if err := r.Run(utils.HttpPort); err != nil {
		panic(err)
	}
}
