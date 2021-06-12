package ginHelper

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
)

type Hello struct {
	Param
	Name string `form:"name" binding:"required"`
}

func (param *Hello) Service(c *gin.Context) {
	param.Ret = getMessage(param.Name)
}

func (param *Hello) Result(c *gin.Context) {
	if param.Err != nil {
		c.String(http.StatusBadRequest, "%s", gin.H{"message": param.Err.Error()})
	} else {
		c.String(http.StatusOK, "%s", param.Ret)
	}
}

type Helper struct{}

func (h *Helper) HelloHandler() (r *Router) {
	return &Router{
		Param:  new(Hello),
		Path:   "/hello",
		Method: "GET",
	}
}

func getMessage(name string) string {
	return "Hello " + name + "!"
}

func TestHelper(t *testing.T) {
	r := gin.New()
	Build(new(Helper), r)
	w := httptest.NewRecorder()
	u := url.URL{Host: "", Path: "hello"}
	query := u.Query()
	name := "Helper"
	query.Add("name", name)
	u.RawQuery = query.Encode()
	req, _ := http.NewRequest("GET", u.String(), nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, getMessage(name), w.Body.String())
}
