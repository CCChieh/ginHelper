package ginHelper

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

type GinRouter interface {
	gin.IRoutes
	BasePath() string
}

func Build(h interface{}, r GinRouter) {

	valueOfh := reflect.ValueOf(h)
	numMethod := valueOfh.NumMethod()
	for i := 0; i < numMethod; i++ {
		rt := valueOfh.Method(i).Call(nil)[0].Interface().(*Router)
		rt.AddHandler(r)
	}
}
