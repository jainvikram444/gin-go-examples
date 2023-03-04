package main

import (
	"net/http"
	"time"

	"gin-example/gin-examples/form"
	_template "gin-example/gin-examples/templates"

	"github.com/gin-gonic/gin"
)

func main() {
	// Disable log's color
	//gin.DisableConsoleColor()

	// Force log's color
	gin.ForceConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
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

	// load html file
	r.LoadHTMLGlob("templates/**/*.tmpl")

	// Load home tempalte as HTML
	r.GET("/template/home", _template.GetHomeTemplate)

	// Load users tempalte as HTML
	r.GET("/template/users", _template.GetUsersTemplate)

	// Quering Data or Post Data
	// http://localhost:8080/formQueryPostData?name=appleboy&address=xyz
	r.GET("/formQueryPostData", form.FormQueryPostData)

	//Middleware: LogWithFormattor will write the logs to gi.DefaultWrtiter
	// By Default gin.DefaultWriter = os.Output

	// Run throgth the below Run() or http configuration
	//r.Run()

	// Or

	// Custom HTTP configuration
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
