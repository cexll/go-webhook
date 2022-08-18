package main

import (
	"github.com/gin-gonic/gin"
	"go_webhook/route"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	route.Load(r)
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
