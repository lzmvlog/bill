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

// 登录
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := new(model.User)
	count, err := util.Engine.Where("username = ? and password = ?", username, password).Count(user)
	if err != nil {
		log.Println(err)
	}
	if count == 1 {
		c.HTML(http.StatusOK, "home.html", gin.H{"username": username})
		c.Redirect(http.StatusMovedPermanently, "/home")
	}
}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{})
}
