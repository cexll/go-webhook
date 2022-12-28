package main

import (
	"flag"
	"fmt"
	"go_webhook/internal/route"

	"github.com/gin-gonic/gin"
)

var port = flag.String("p", "8080", "the is run port")

func main() {
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	route.Load(r)
	fmt.Printf("go webhook is running... 0.0.0.0:%v", *port)
	err := r.Run(fmt.Sprintf(":%v", *port))

	if err != nil {
		panic(err)
	}
}
