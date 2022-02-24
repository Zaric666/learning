package main

import (
	"github.com/Zaric666/learning/gin/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.InitRouter(r)

	r.Run(":8183")
}
