package ginHelper

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHelper(t *testing.T) {
	router := gin.Default()
	r := router.Group("api")
	h := NewWithSwagger(&SwaggerInfo{
		Description: "swagger test page",
		Title:       "Swagger Test Page",
		Version:     "0.0.1",
		ContactInfoProps: ContactInfoProps{
			Name:  "zzj",
			URL:   "https://zzj.cool",
			Email: "email@zzj.cool",
		},
	}, r)
	h.Add(testGroup, r)
	router.Run(":8888")
}

var testGroup = &GroupRouter{
	Path: "test",
	Name: "Mytest",
	Routers: []*Router{
		{
			Param:  new(testBodyParam),
			Path:   "/hello/:id",
			Method: "POST",
		}},
}

type testBodyParam struct {
	BaseParam `json:"-"`
	Foo       string `binding:"required"`
	FooName   string `json:"fName" binding:"required"`
	FooInt    int    `binding:"required"`
	FooStruct
	FooStruct2 FooStruct
	FooStruct3 *FooStruct
}

type FooStruct struct {
	FooA string `binding:"required"`
	FooB bool   `binding:"required"`
}

func (param *testBodyParam) Handler(c *gin.Context) (Data, error) {
	return param, nil
}
