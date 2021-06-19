package ginHelper

import (
	"fmt"
	"path"
	"reflect"
)

type routerView map[string]map[string]*Router

type Helper struct {
	routers routerView
	Swagger *Swagger
}

func New() *Helper {
	return &Helper{routers: routerView{}}
}

func NewWithSwagger(r GinRouter) *Helper {
	swg := &Swagger{
		Router: r.Group("swagger"),
		SwaggerInfo: &SwaggerInfo{
			BasePath:    r.BasePath(),
			Description: "Swagger test",
			Title:       "GinHelper Swagger",
		},
	}
	swg.Init()
	return &Helper{routers: routerView{}, Swagger: swg}
}

func (h *Helper) Add(helper interface{}, r GinRouter) {
	valueOfh := reflect.ValueOf(helper)
	elemName := reflect.TypeOf(helper).Elem().Name()
	numMethod := valueOfh.NumMethod()

	for i := 0; i < numMethod; i++ {
		rt := valueOfh.Method(i).Call(nil)[0].Interface().(*Router)
		rt.AddHandler(r)
		h.addPath(rt, r, elemName)
	}
}

func (h *Helper) addPath(rt *Router, r GinRouter, elemName string) {
	if h.Swagger == nil {
		return
	}

	typeOf := reflect.TypeOf(rt.Param).Elem()
	for i := 0; i < typeOf.NumField(); i++ {
		fmt.Println(typeOf.Field(i).Name)
	}

	apiPath := path.Join(h.cleanPath(h.Swagger.BasePath, r.BasePath()), rt.Path)
	h.Swagger.AddPath(&SwaggerPath{
		Path:   apiPath,
		Method: rt.Method,
		Tags:   []string{elemName},
		Param:  rt.Param,
	})
	_, ok := h.routers[apiPath]
	if !ok {
		h.routers[apiPath] = map[string]*Router{}
	}
	h.routers[apiPath][rt.Method] = rt
}

func (h *Helper) cleanPath(basePath, path string) string {
	for i := 0; i < len(path); i++ {
		if i >= len(basePath) || basePath[i] != path[i] {
			return path[i:]
		}
	}
	return ""
}

func (h *Helper) View() routerView { return h.routers }
