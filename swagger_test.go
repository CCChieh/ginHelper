package ginHelper

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSwagger(t *testing.T) {
	swaggerR := gin.Default()
	swagger(swaggerR.Group("swagger"))
	// TODO解决测试
	swaggerR.Run(":8888")
}
