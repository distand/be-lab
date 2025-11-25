package router

import (
	"be-lab/middleware"
	"be-lab/server"
	"github.com/gin-gonic/gin"
)

func SetupRouter(s *server.Server) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), middleware.Log())

	// 无需登录态
	r.POST("/api/login", s.Login) // 登录
	// API
	a := r.Group("/api/:openid", s.CheckLogin)
	a.GET("/index", s.Index) // 首页
	// 用户相关
	u := a.Group("/user")
	u.GET("/info", s.UserInfo)  // 用户信息
	u.POST("/save", s.UserSave) // 保存用户
	// 仪器相关
	d := a.Group("/device")
	d.GET("/list", s.DeviceList) // 仪器列表
	// 预约相关
	b := a.Group("/booking")
	b.GET("/list", s.BookingList)  //	预约列表
	b.POST("/save", s.BookingSave) // 保存预约

	//管理员
	c := a.Group("/admin", s.CheckAdmin)
	c.GET("/device/save", s.DeviceSave) // 保存仪器
	c.GET("/user/list", s.UserList)     // 用户列表
	c.GET("/user/save", s.UserSave)     // 保存用户
	return r
}
