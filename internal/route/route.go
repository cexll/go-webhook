package route

import (
	"github.com/gin-gonic/gin"
	"go_webhook/internal/controllers"
)

func Load(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello world")
	})
	r.POST("/github", controllers.ParseGithub)
	r.POST("/gitee", controllers.ParseGitee)
}
