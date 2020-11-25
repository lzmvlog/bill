package routers

import (
	"bill/app/api"
	"bill/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	router := gin.New()
	// 初始化数据连接
	util.Init()
	// 加载模板
	router.LoadHTMLGlob("resources/templates/*/*")
	// 加载静态资源文件
	router.StaticFS("/resources/static/", http.Dir("resources/static/"))
	// 路由路由
	router.GET("/login", api.Login)
	return router
}
