package templates

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHomeTemplate(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index.tmpl", gin.H{
		"title": "Gin -  home Template",
	})
}

func GetUsersTemplate(c *gin.Context) {
	c.HTML(http.StatusOK, "users/users.tmpl", gin.H{
		"title": "Gin -  User Template",
	})
}
