package ginHelper

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSwagger(t *testing.T) {
	swaggerR := gin.Default()
	swg := &Swagger{
		Router:   swaggerR.Group("swagger"),
		BasePath: "/api",
		SwaggerInfo: &SwaggerInfo{
			Description: "Swagger test",
			Title:       "GinHelper Swagger",
		},
	}
	swg.Init()
	swg.AddPath(&SwaggerApi{
		Path:   "/testsadfdsdd/id",
		Method: "GET",
		Tags:   []string{"dfd"},
	})
	swg.AddPath(&SwaggerApi{
		Path:   "testsadfdsdd",
		Method: "POST",
	})
	// TODO解决测试
	swaggerR.Run(":8888")
}
