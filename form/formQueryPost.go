package form

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func FormQueryPostData(c *gin.Context) {
	var person Person
	err := c.ShouldBind(&person)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"Person": person,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})
	}
}
