package ginHelper

import (
	"github.com/gin-gonic/gin"
	"strings"
)

type Router struct {
	Path     string
	Method   string
	Handlers []gin.HandlerFunc
}

func (rt *Router) AddHandler(r gin.IRoutes) {
	switch strings.ToUpper(rt.Method) {
	case "GET":
		r.GET(rt.Path, rt.Handlers...)
	case "POST":
		r.POST(rt.Path, rt.Handlers...)
	case "PUT":
		r.PUT(rt.Path, rt.Handlers...)
	case "PATCH":
		r.PATCH(rt.Path, rt.Handlers...)
	case "HEAD":
		r.HEAD(rt.Path, rt.Handlers...)
	case "OPTIONS":
		r.OPTIONS(rt.Path, rt.Handlers...)
	case "DELETE":
		r.DELETE(rt.Path, rt.Handlers...)
	case "ANY":
		r.Any(rt.Path, rt.Handlers...)
	default:
		panic("Method: " + rt.Method + " is error")
	}
}
