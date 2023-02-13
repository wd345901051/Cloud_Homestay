package routers

import (
	"cloud_homestay/helper"
	"cloud_homestay/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(helper.EnableCookieSession())
	// 路由规则

	r.Static("/static", "static")

	// 自定义路由方法
	// r.SetFuncMap(template.FuncMap{
	// 	"deltime": helper.Deltime,
	// })
	r.LoadHTMLGlob("templates/**/*")
	// 路由方法
	// 登录页面
	r.GET("/login", service.LoginPage)
	// 注册业务
	r.POST("/userlogup", service.UserLogup)
	// 登录业务
	r.POST("/userlogin", service.UserLogin)
	// 首页页面
	r.GET("/", helper.GetUserSessionHandler, service.Home)
	// 首页页面
	r.GET("/home", helper.GetUserSessionHandler, service.Home)

	// 房间列表
	r.GET("/room", helper.GetUserSessionHandler, service.RoomList)

	// 房间排行分类
	// 默认排行
	r.GET("/room1", service.Room1)
	// 销量排行
	r.GET("/room2", service.Room2)
	// 价格排行
	r.GET("/room3", service.Room3)
	// 房间信息
	r.GET("/roominfo", helper.GetUserSessionHandler, service.CheckLogin, service.RoomInfo)
	// 订单信息
	r.POST("/order", helper.GetUserSessionHandler, service.OrderPush)

	// 评论信息
	r.POST("/comment", service.Comment)
	// 个人中心
	r.GET("/personal", helper.GetUserSessionHandler, service.CheckLogin, service.Personal)

	// admin业务

	admin := r.Group("/admin", helper.GetUserSessionHandler, service.CheckAdmin)
	{
		// 管理房间页面
		admin.GET("/room", service.Room)
		// 管理订单页面
		admin.GET("/order", service.Order)
		// 用户管理页面
		admin.GET("/user", service.User)
		// 评论管理
		admin.GET("/message", service.Message)
		// 货物管理
		admin.GET("/goods", service.Goods)
		// 设备管理
		admin.GET("/equip", service.Equip)
		// 用户统计
		admin.GET("/userct", service.UserCT)

	}

	// 业务处理

	// 关注网站
	r.POST("/attention", helper.GetUserSessionHandler, service.CheckLogin, service.Attention)

	// 退出登录
	r.GET("/rmsession", service.Rmsession)

	// banuser
	r.POST("/banuser", service.BanUser)

	// 删除房间
	r.POST("/deleteroom", service.DeleteRoom)

	// 添加房间
	r.POST("/roomadd", service.RoomAdd)

	// 删除物品
	r.POST("/deletegood", service.DeleteGood)

	// 添加物品
	r.POST("/goodadd", service.GoodAdd)

	// 删除评论
	r.POST("/deletcomment", service.DeleteComment)

	return r
}
