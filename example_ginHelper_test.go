package ginHelper

import (
	"testing"

	"github.com/gin-gonic/gin"
)

var exGroup = &GroupRouter{
	Path: "example",
	Name: "Mytest",
	Routes: []*Route{
		{
			Param:  new(exParam),
			Path:   "/foo/:id",
			Method: "POST",
		}},
}

type exParam struct {
	BaseParam
	Foo  string `json:"foo"`
	Id   string `uri:"id"`
	Auth string `header:"auth" json:"-"`
}

func (param *exParam) Handler(c *gin.Context) (data Data, err error) {
	return param, nil
}

func ExampleNewWithSwagger() {
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
	h.Add(exGroup, r)
	_ = router.Run(":12321")
}

func TestExample(t *testing.T) {
	ExampleNewWithSwagger()
}
