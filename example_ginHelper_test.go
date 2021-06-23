package ginHelper

import (
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
	_ = router.Run(":8888")
}

// func TestExample(t *testing.T) {
// 	ExampleNewWithSwagger()
// }
