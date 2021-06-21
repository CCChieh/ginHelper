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

func (h *Helper) Add(gh *GroupRouter, r GinRouter) {
	r = r.Group(gh.Path)
	for _, rt := range gh.Routers {
		rt.AddHandler(r)
		h.addPath(rt, r, gh.Name)
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
	h.Swagger.AddPath(&SwaggerApi{
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
