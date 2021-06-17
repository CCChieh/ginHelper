package ginHelper

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSwagger(t *testing.T) {
	swaggerR := gin.Default()
	swg := &swagger{
		SwaggerInfo: &SwaggerInfo{
			BasePath:    "/api",
			Description: "Swagger test",
			Title:       "GinHelper Swagger",
		},
	}
	swg.mount(swaggerR.Group("swagger"))
	// TODO解决测试
	swaggerR.Run(":8888")
}
