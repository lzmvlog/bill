package api

import (
	"bill/app/model"
	"bill/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 登录页面
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// 主页面
func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{})
}

// 登录
func Login(c *gin.Context) {
	user := new(model.User)
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	count, err := util.Engine.Count(user)
	if count != 1 {
		log.Println(err)
		c.JSON(200, count)
	}

	if count == 1 {
		c.JSON(http.StatusOK, count)
	}
}
