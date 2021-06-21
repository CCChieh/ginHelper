package ginHelper

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHelper(t *testing.T) {
	router := gin.Default()
	r := router.Group("api")
	h := NewWithSwagger(r)
	h.Add(GroupR, r)
	router.Run(":8888")
}

var GroupR = &GroupRouter{
	Path: "my",
	Name: "Mytest",
	Routers: []*Router{
		{
			Param:  new(HelloParam),
			Path:   "/hello/:kkk",
			Method: "POST",
		}},
}

type HelloParam struct {
	BaseParam
	Vparam
	Vparam1 Vparam
	Vparam2 *Vparam
	Foo     string `json:"foo4"`
	Va
	Ch    byte
	Arr   []string
	Arr2  []int
	Arr3  []Vparam
	Int   int
	Float float32
	Bool  bool
	// Bar string `json:"bar" binding:"required"`
}

type Va bool

type Vparam struct {
	Bar  string `json:"bar" binding:"required"`
	Bare bool   `json:"bare" binding:"required"`
}

func (param *HelloParam) Handler(c *gin.Context) (Data, error) {
	return getMessage(param.Foo) + getMessage(param.Bar), nil
}
