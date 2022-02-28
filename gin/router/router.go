package router

import (
	"github.com/Zaric666/learning/gin/middleware/logger"
	v1 "github.com/Zaric666/learning/gin/router/v1"
	"github.com/Zaric666/learning/gin/validator/user"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
)

func InitRouter(r *gin.Engine) {

	// test middleware
	r.Use(logger.Logger())
	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("laosan").(int)
		// it would print: 123456
		log.Println(example)
	})

	// v1 api
	v1Api := r.Group("/v1")
	{
		v1Api.GET("/users", v1.User{}.Index)
	}

	// 绑定验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("NameValid", user.NameValid)
	}
}
