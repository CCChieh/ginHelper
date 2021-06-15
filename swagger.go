package ginHelper

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed swagger
var swaggerFS embed.FS

func Swagger(path string, r GinRouter) {
	fads, _ := fs.Sub(swaggerFS, "swagger")
	_ = r.StaticFS(path, http.FS(fads))
}
