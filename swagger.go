package ginHelper

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed swagger
var swaggerFS embed.FS

func Swagger(path string, r *gin.Engine) {
	fads, _ := fs.Sub(swaggerFS, "swagger")
	r.StaticFS(path, http.FS(fads))
}
