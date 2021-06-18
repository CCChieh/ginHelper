package ginHelper

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHelper(t *testing.T) {
	router := gin.Default()
	r := router.Group("api")
	h := NewWithSwagger(r)
	h.Add(new(HelloNewHelper), r)
	router.Run(":8888")
}

type HelloNewHelper struct{}

func (h *HelloNewHelper) HelloHandler() (r *Router) {
	return &Router{
		Param:  new(HelloParam),
		Path:   "/hello",
		Method: "GET",
	}
}

type HelloParam struct {
	BaseParam
	Foo string `form:"foo" binding:"required"`
}

func (param *HelloParam) Service(c *gin.Context) (Data, error) {
	return getMessage(param.Foo), nil
}
