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
	Vparam
	Foo string `form:"foo"`
	// Bar string `json:"bar" binding:"required"`
}

type Vparam struct {
	Bar string `form:"bar" binding:"required"`
}

func (param *HelloParam) Service(c *gin.Context) (Data, error) {
	return getMessage(param.Foo) + getMessage(param.Bar), nil
}
