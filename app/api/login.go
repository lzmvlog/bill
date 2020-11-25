package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
