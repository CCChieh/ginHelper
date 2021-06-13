package ginHelper

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
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

type HelloHelper struct{}

func (h *HelloHelper) HelloHandler() (r *Router) {
	return &Router{
		Param:  new(Hello),
		Path:   "/hello1",
		Method: "GET",
	}
}

func Help2(c *gin.Context) {
	p := &Hello{}
	if err := c.ShouldBind(p); err != nil {
		p.Err = err
	}
	if p.Err != nil {
		c.String(http.StatusBadRequest, "%s", gin.H{"message": p.Err.Error()})
	} else {
		c.String(http.StatusOK, "%s", getMessage(p.Name))
	}
}

func getMessage(name string) string {
	return "Hello " + name + "!"
}

var r0 *gin.Engine
var r1 *gin.Engine
var r2 *gin.Engine

func TestMain(m *testing.M) {
	gin.SetMode(gin.ReleaseMode)

	r0 = gin.New()
	h := New()
	h.Add(new(HelloHelper), r0)
	h.Add(new(HelloHelper), r0.Group("api"))
	fmt.Println(h.View())
	for path, v := range h.View() {
		for method, router := range v {
			fmt.Println(path, method, router.Param)
		}
	}

	r1 = gin.New()
	Build(new(HelloHelper), r1)

	r2 = gin.New()
	r2.GET("/hello2", Help2)

	code := m.Run()
	os.Exit(code)
}

func TestGinHelper(t *testing.T) {
	name := "Helper"
	w := httptest.NewRecorder()
	u := url.URL{Host: "", Path: "hello1", RawQuery: url.Values{"name": []string{name}}.Encode()}

	req, _ := http.NewRequest("GET", u.String(), nil)
	r1.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, getMessage(name), w.Body.String())
}

func TestGin(t *testing.T) {
	name := "Helper"
	w := httptest.NewRecorder()
	u := url.URL{Host: "", Path: "hello2", RawQuery: url.Values{"name": []string{name}}.Encode()}

	req, _ := http.NewRequest("GET", u.String(), nil)
	r2.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, getMessage(name), w.Body.String())
}

func BenchmarkGinHelper(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		name := "Helper"
		w := httptest.NewRecorder()
		u := url.URL{Host: "", Path: "hello1", RawQuery: url.Values{"name": []string{name}}.Encode()}

		req, _ := http.NewRequest("GET", u.String(), nil)
		r1.ServeHTTP(w, req)
	}
}

func BenchmarkGin(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		name := "Helper"
		w := httptest.NewRecorder()
		u := url.URL{Host: "", Path: "hello2", RawQuery: url.Values{"name": []string{name}}.Encode()}

		req, _ := http.NewRequest("GET", u.String(), nil)
		r2.ServeHTTP(w, req)
	}
}
