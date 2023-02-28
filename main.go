package main

import (
	"net/http"

	"gin-example/gin-examples/form"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// ping for the health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// asciiOnlyJSON for the return ascii code for the non-ascii char
	r.GET("/asciiOnlyJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br/>",
		}

		c.AsciiJSON(http.StatusOK, data)
	})

	// Form data for the return form data of Struct of form
	// http://localhost:8080/formGetB?field_a=hello&field_b=world
	r.GET("/formGetB", form.GetDataB)

	// http://localhost:8080/formGetC?field_a=hello&field_c=world
	r.GET("/formGetC", form.GetDataC)

	// http://localhost:8080/formGetD?field_x=hello&field_d=world
	r.GET("/formGetD", form.GetDataD)

	//browse html file <Relative path>/public/formColors.html
	r.POST("/formColors", form.FormHandler)

	r.Run()
}
