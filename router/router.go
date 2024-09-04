package router

import (
	"backend/controller"
	"backend/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	// 解决跨域问题
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                       // 允许的来源
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},            // 允许的请求方法
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"}, // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"},                          // 暴露的响应头
		AllowCredentials: true,                                                // 允许传递凭据（例如 Cookie）
		MaxAge:           12 * time.Hour,                                      // 预检请求的有效期
	}))

	// 设置静态文件目录
	r.Static("/static", "./dist/static")
	r.LoadHTMLGlob("dist/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// 测试 GET 请求
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 注册
	r.POST("/register", controller.Register)

	// 登录
	r.POST("/login", controller.Login)

	// 登出
	r.POST("/logout", controller.Logout)

	// 大模型上链
	r.POST("/uplink", middleware.JWTAuthMiddleware(), controller.Uplink)

	// 获取大模型的上链信息
	r.POST("/getModelInfo", controller.GetModelInfo)

	// 获取用户的大模型ID列表
	r.POST("/getModelList", middleware.JWTAuthMiddleware(), controller.GetModelList)

	// 获取所有的大模型信息
	r.POST("/getAllModelInfo", middleware.JWTAuthMiddleware(), controller.GetAllModelInfo)

	// 获取大模型上链历史(溯源)
	r.POST("/getModelHistory", middleware.JWTAuthMiddleware(), controller.GetModelHistory)

	// 查询用户的积分
	r.POST("/getUserPoints", middleware.JWTAuthMiddleware(), controller.GetUserPoints)

	// 上传模型
	r.POST("/upload-model", middleware.JWTAuthMiddleware(), controller.UploadModel)

	// 下载模型
	r.POST("/download-model", middleware.JWTAuthMiddleware(), controller.DownloadModel)

	return r
}
