package ginHelper

import (
	"embed"
	"io/fs"
	"net/http"
	"path"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/spec"
)

//go:embed swagger
var swaggerFS embed.FS

type SwaggerInfo struct {
	BasePath    string
	Description string
	Title       string
}

type SwaggerPath struct {
	Path        string
	Method      string
	Description string
	Summary     string
	Tags        []string
	Param       interface{}
}

type Swagger struct {
	Router GinRouter
	*SwaggerInfo
	Spec *spec.Swagger
}

func (s *Swagger) Init() {
	if s.Spec == nil {
		s.genSwaggerJson()
	}
	// s.AddPath("/testsadfds", "GET")
	s.Router.Use(func(c *gin.Context) {
		if path.Base(c.Request.URL.Path) == "swagger.json" {
			c.Writer.Header().Set("content-type", "text/json")
			s.Spec.SwaggerProps.Host = c.Request.Host
			c.JSON(200, s.Spec)
			c.Abort()
		}
		c.Next()
	})
	swaggerDir, _ := fs.Sub(swaggerFS, "swagger")
	_ = s.Router.StaticFS("", http.FS(swaggerDir))
}

func (s *Swagger) AddPath(sp *SwaggerPath) {
	sp.Path = path.Join("/", sp.Path)
	sp.Path = path.Clean(sp.Path)
	_, ok := s.Spec.Paths.Paths[sp.Path]
	if !ok {
		s.Spec.Paths.Paths[sp.Path] = spec.PathItem{
			PathItemProps: spec.PathItemProps{},
		}
	}
	test := &spec.Schema{}
	test.SetProperty("dddsf", *spec.StringProperty())

	operation := &spec.Operation{
		VendorExtensible: spec.VendorExtensible{},
		OperationProps: spec.OperationProps{
			Description: sp.Description,
			Tags:        sp.Tags,
			Summary:     sp.Summary,
			Parameters:  s.parameters(sp),
			Responses:   &spec.Responses{},
		},
	}
	temp := s.Spec.Paths.Paths[sp.Path]
	switch strings.ToUpper(sp.Method) {
	case "GET":
		temp.Get = operation
	case "POST":
		temp.Post = operation
	case "PUT":
		temp.Put = operation
	case "PATCH":
		temp.Patch = operation
	case "HEAD":
		temp.Head = operation
	case "OPTIONS":
		temp.Options = operation
	case "DELETE":
		temp.Delete = operation
	case "ANY":
		temp.Get = operation
		temp.Post = operation
		temp.Put = operation
		temp.Patch = operation
		temp.Head = operation
		temp.Options = operation
		temp.Delete = operation
	default:

	}
	s.Spec.Paths.Paths[sp.Path] = temp
	schema := &spec.Schema{}
	schema2 := &spec.Schema{}
	schema.
		SetProperty("dfd", *spec.StringProperty()).
		SetProperty("dfdee", *spec.Int64Property()).
		SetProperty("aaa", spec.Schema{SchemaProps: spec.SchemaProps{Type: []string{"object"}}}).
		SetProperty("ikj", *spec.ComposedSchema(*schema2))

	s.Spec.Definitions["Pet"] = *schema
}

func (s *Swagger) genSwaggerJson() {
	s.Spec = &spec.Swagger{
		SwaggerProps: spec.SwaggerProps{
			Swagger: "2.0",
			Info: &spec.Info{
				InfoProps: spec.InfoProps{
					Description: s.SwaggerInfo.Description,
					Title:       s.SwaggerInfo.Title,
					Contact: &spec.ContactInfo{
						ContactInfoProps: spec.ContactInfoProps{
							Name:  "zzj",
							URL:   "https://zzj.cool",
							Email: "email@zzj.cool",
						},
					},
					Version: "0.0.1",
				},
			},
			BasePath: s.SwaggerInfo.BasePath,
			Paths: &spec.Paths{
				Paths: map[string]spec.PathItem{},
			},
			Definitions: spec.Definitions{},
		},
	}
}

func (s *Swagger) parameters(sp *SwaggerPath) (params []spec.Parameter) {
	param := sp.Param
	method := sp.Method
	params = append(params, pathParams(sp)...)

	if param == nil {
		return nil
	}
	if method == http.MethodGet {
		params = append(params, queryParams(reflect.TypeOf(param))...)
	}
	params = append(params, *spec.BodyParam("body", JsonSchemas(reflect.TypeOf(param))))
	return params
}
