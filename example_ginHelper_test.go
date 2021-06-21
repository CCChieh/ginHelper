package ginHelper

import (
	"github.com/gin-gonic/gin"
)

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
	h.Add(testGroup, r)
	router.Run(":8888")
}
