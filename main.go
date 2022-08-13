package main

import (
	"github.com/gin-gonic/gin"
	"go_webhook/controllers"
)

func main() {
	gin.SetMode("release")
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello world")
	})
	r.POST("/github", controllers.ParseGithub)
	r.POST("/gitee", controllers.ParseGitee)
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
