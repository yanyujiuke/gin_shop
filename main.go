package main

import (
	"ginShop/common"
	"ginShop/routers"
	"github.com/gin-gonic/gin"
	"html/template"
)

func main() {
	r := gin.Default()

	// 自定义模板函数  注意要把这个函数放在加载模板前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": common.UnixToTime,
	})

	// 加载模板
	r.LoadHTMLGlob("views/**/**/*")

	// 配置静态web目录
	r.Static("/static", "./static")

	// 注册路由
	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)

	r.Run()
}
