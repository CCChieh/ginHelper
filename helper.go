package ginHelper

import (
	"path"
	"reflect"
)

type routerView map[string]map[string]*Router

type Helper struct {
	routers routerView
}

func New() *Helper {
	return &Helper{routers: routerView{}}
}

func (h *Helper) Add(helper interface{}, r GinRouter) {
	valueOfh := reflect.ValueOf(helper)
	numMethod := valueOfh.NumMethod()
	for i := 0; i < numMethod; i++ {
		rt := valueOfh.Method(i).Call(nil)[0].Interface().(*Router)
		rt.AddHandler(r)
		path := path.Join(r.BasePath(), rt.Path)
		_, ok := h.routers[path]
		if !ok {
			h.routers[path] = map[string]*Router{}
		}
		h.routers[path][rt.Method] = rt
	}
}

func (h *Helper) Swagger(r GinRouter) {
	swg := &swagger{}
	swg.mount(r.Group("swagger"))
}

func (h *Helper) View() routerView { return h.routers }
