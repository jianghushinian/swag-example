package main

import (
	"github.com/gin-gonic/gin"

	"github.com/jianghushinian/swag-example/internal/api"
)

func main() {
	// 代码中设置参数会覆盖 swag-example/cmd/docs.go 中的注释
	// docs.SwaggerInfo.Title = "codes set title"

	r := gin.Default()
	api.InitController(r)
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
