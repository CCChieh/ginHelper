package ginHelper

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSwagger(t *testing.T) {
	swaggerR := gin.Default()
	swg := &Swagger{
		Router: swaggerR.Group("swagger"),
		SwaggerInfo: &SwaggerInfo{
			BasePath:    "/api",
			Description: "Swagger test",
			Title:       "GinHelper Swagger",
		},
	}
	swg.Init()
	swg.AddPath("/testsadfdsdd", "GET")
	swg.AddPath("/testsadfdsdd", "POST")
	// TODO解决测试
	swaggerR.Run(":8888")
}
