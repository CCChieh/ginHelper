package ginHelper

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

func Build(h interface{}, r gin.IRoutes) {
	valueOfh := reflect.ValueOf(h)
	numMethod := valueOfh.NumMethod()
	for i := 0; i < numMethod; i++ {
		rt := valueOfh.Method(i).Call(nil)[0].Interface().(*Router)
		rt.AddHandler(r)
	}
}
