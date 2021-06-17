package ginHelper

import (
	"embed"
	"io/fs"
	"net/http"
	"path"

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

type swagger struct {
	*SwaggerInfo
	SwaggerConf *spec.Swagger
}

func (s *swagger) mount(r GinRouter) {
	if s.SwaggerConf == nil {
		s.genSwaggerJson()
	}
	r.Use(func(c *gin.Context) {
		if path.Base(c.Request.URL.Path) == "swagger.json" {
			c.Writer.Header().Set("content-type", "text/json")
			s.SwaggerConf.SwaggerProps.Host = c.Request.Host
			c.JSON(200, s.SwaggerConf)
			c.Abort()
		}
		c.Next()
	})
	swaggerDir, _ := fs.Sub(swaggerFS, "swagger")
	_ = r.StaticFS("", http.FS(swaggerDir))
}

func (s *swagger) genSwaggerJson() {
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
				VendorExtensible: spec.VendorExtensible{},
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
				}},
			},
			Definitions:         map[string]spec.Schema{},
			Parameters:          map[string]spec.Parameter{},
			Responses:           map[string]spec.Response{},
			SecurityDefinitions: map[string]*spec.SecurityScheme{},
			Security:            []map[string][]string{},
			Tags:                []spec.Tag{},
			ExternalDocs:        &spec.ExternalDocumentation{},
		},
	}
}
