package ginHelper

import (
	"path"
)

type routerView map[string]map[string]*Route

type Helper struct {
	routers routerView
	Swagger *Swagger
}

func New() *Helper {
	return &Helper{routers: routerView{}}
}

func NewWithSwagger(swaggerInfo *SwaggerInfo, r GinRouter) *Helper {
	swg := &Swagger{
		Router:      r.Group("swagger"),
		BasePath:    r.BasePath(),
		SwaggerInfo: swaggerInfo,
	}
	swg.Init()
	return &Helper{routers: routerView{}, Swagger: swg}
}

func (h *Helper) Add(gh *GroupRouter, r GinRouter) {
	r = r.Group(gh.Path)
	h.Swagger.AddTag(gh.Name, gh.Description)
	for _, rt := range gh.Routes {
		rt.AddHandler(r)
		h.addPath(rt, r, gh.Name)
	}
}

func (h *Helper) addPath(rt *Route, r GinRouter, elemName string) {
	if h.Swagger == nil {
		return
	}

	apiPath := path.Join(h.cleanPath(h.Swagger.BasePath, r.BasePath()), rt.Path)
	h.Swagger.AddPath(&SwaggerApi{
		Summary: rt.Summary,
		Path:    apiPath,
		Method:  rt.Method,
		Tags:    []string{elemName},
		Param:   rt.Param,
	})
	_, ok := h.routers[apiPath]
	if !ok {
		h.routers[apiPath] = map[string]*Route{}
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
