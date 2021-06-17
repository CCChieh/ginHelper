package ginHelper

import (
	"embed"
	"io/fs"
	"net/http"
	"path"
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

type Swagger struct {
	Router GinRouter
	*SwaggerInfo
	SwaggerConf *spec.Swagger
}

func (s *Swagger) Init() {
	if s.SwaggerConf == nil {
		s.genSwaggerJson()
	}
	// s.AddPath("/testsadfds", "GET")
	s.Router.Use(func(c *gin.Context) {
		if path.Base(c.Request.URL.Path) == "swagger.json" {
			c.Writer.Header().Set("content-type", "text/json")
			s.SwaggerConf.SwaggerProps.Host = c.Request.Host
			c.JSON(200, s.SwaggerConf)
			c.Abort()
		}
		c.Next()
	})
	swaggerDir, _ := fs.Sub(swaggerFS, "swagger")
	_ = s.Router.StaticFS("", http.FS(swaggerDir))
}

func (s *Swagger) AddPath(path, method string) {
	_, ok := s.SwaggerConf.Paths.Paths[path]
	if !ok {
		s.SwaggerConf.Paths.Paths[path] = spec.PathItem{
			PathItemProps: spec.PathItemProps{},
		}
	}
	operation := &spec.Operation{
		VendorExtensible: spec.VendorExtensible{},
		OperationProps: spec.OperationProps{
			Description:  "ddd",
			Consumes:     []string{},
			Produces:     []string{},
			Schemes:      []string{},
			Tags:         []string{"Dfd"},
			Summary:      "",
			ExternalDocs: &spec.ExternalDocumentation{},
			ID:           "",
			Deprecated:   false,
			Security:     []map[string][]string{},
			Parameters:   []spec.Parameter{},
			Responses:    &spec.Responses{},
		},
	}
	temp := s.SwaggerConf.Paths.Paths[path]
	switch strings.ToUpper(method) {
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
	s.SwaggerConf.Paths.Paths[path] = temp
}

func (s *Swagger) genSwaggerJson() {
	s.SwaggerConf = &spec.Swagger{
		VendorExtensible: spec.VendorExtensible{
			Extensions: map[string]interface{}{},
		},
		SwaggerProps: spec.SwaggerProps{
			ID:       "",
			Consumes: []string{},
			Produces: []string{},
			Schemes:  []string{},
			Swagger:  "2.0",
			Info: &spec.Info{
				InfoProps: spec.InfoProps{
					Description:    s.SwaggerInfo.Description,
					Title:          s.SwaggerInfo.Title,
					TermsOfService: "",
					Contact: &spec.ContactInfo{
						ContactInfoProps: spec.ContactInfoProps{
							Name:  "zzj",
							URL:   "https://zzj.cool",
							Email: "email@zzj.cool",
						},
					},
					License: &spec.License{},
					Version: "0.0.1",
				},
			},
			BasePath: s.SwaggerInfo.BasePath,
			Paths: &spec.Paths{
				Paths: map[string]spec.PathItem{"/pet/{petId}/uploadImage": {
					Refable: spec.Refable{
						Ref: spec.Ref{},
					},
					VendorExtensible: spec.VendorExtensible{
						Extensions: map[string]interface{}{},
					},
					PathItemProps: spec.PathItemProps{
						Get: &spec.Operation{
							VendorExtensible: spec.VendorExtensible{},
							OperationProps: spec.OperationProps{
								Description:  "f444",
								Consumes:     []string{},
								Produces:     []string{},
								Schemes:      []string{},
								Tags:         []string{"hello"},
								Summary:      "sss",
								ExternalDocs: &spec.ExternalDocumentation{},
								ID:           "",
								Deprecated:   false,
								Security:     []map[string][]string{},
								Parameters: []spec.Parameter{{
									Refable:           spec.Refable{},
									CommonValidations: spec.CommonValidations{},
									SimpleSchema:      spec.SimpleSchema{},
									VendorExtensible:  spec.VendorExtensible{},
									ParamProps: spec.ParamProps{
										Description:     "",
										Name:            "test",
										In:              "query",
										Required:        false,
										Schema:          &spec.Schema{},
										AllowEmptyValue: false,
									},
								}},
								Responses: &spec.Responses{},
							},
						},
					},
				},
				},
			},

			Responses:           map[string]spec.Response{},
			SecurityDefinitions: map[string]*spec.SecurityScheme{},
			Security:            []map[string][]string{},
			Tags:                []spec.Tag{},
			ExternalDocs:        &spec.ExternalDocumentation{},
		},
	}
}
