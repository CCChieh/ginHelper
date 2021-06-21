package ginHelper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

var testGroup = &GroupRouter{
	Path: "api",
	Name: "Mytest",
	Routers: []*Router{
		{
			Param:  new(testBodyParam),
			Path:   "/hello/:id",
			Method: "POST",
		}},
}

type FooStruct struct {
	FooA string `binding:"required" `
	FooB *bool  `binding:"required"`
}

type testBodyParam struct {
	BaseParam `json:"-"`
	Foo       string `binding:"required"`
	FooName   string `json:"fName" binding:"required"`
	FooInt    int    `binding:"required"`
	FooIgnore string `json:"-"`
	FooStruct
	FooStruct2 FooStruct
	FooStruct3 *FooStruct
}

func (param *testBodyParam) Handler(c *gin.Context) (Data, error) {
	return param, nil
}

func TestHelperWithSwagger(t *testing.T) {
	router := gin.Default()
	r := router.Group("api")
	h := NewWithSwagger(&SwaggerInfo{
		Description: "swagger test page",
		Title:       "Swagger Test Page",
		Version:     "0.0.1",
		ContactInfoProps: ContactInfoProps{
			Name:  "zzj",
			URL:   "https://zzj.cool",
			Email: "email@zzj.cool",
		},
	}, r)
	h.Add(testGroup, r)
	router.Run(":8888")
}

func TestNew(t *testing.T) {
	False := false
	tests := []struct {
		name   string
		input  Parameter
		want   Parameter
		path   string
		method string
	}{
		{
			name: "POST",
			input: &testBodyParam{
				Foo:       "bar",
				FooName:   "bar",
				FooInt:    9,
				FooIgnore: "bar",
				FooStruct: FooStruct{
					FooA: "bar",
					FooB: &False,
				},
				FooStruct2: FooStruct{
					FooA: "bar",
					FooB: &False,
				},
				FooStruct3: &FooStruct{
					FooA: "bar",
					FooB: &False,
				},
			},
			want: &testBodyParam{
				Foo:     "bar",
				FooName: "bar",
				FooInt:  9,
				FooStruct: FooStruct{
					FooA: "bar",
					FooB: &False,
				},
				FooStruct2: FooStruct{
					FooA: "bar",
					FooB: &False,
				},
				FooStruct3: &FooStruct{
					FooA: "bar",
					FooB: &False,
				},
			},
			path:   "/api/hello/fooId",
			method: "POST",
		},
	}
	router := gin.Default()
	h := New()
	h.Add(testGroup, router)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			jsonParam, err := json.Marshal(tt.input)

			if err != nil {
				t.Fatalf("Json Marshal fail for: %v", tt.input)
			}
			fmt.Println(string(jsonParam))
			req, _ := http.NewRequest(tt.method, tt.path, bytes.NewBuffer(jsonParam))
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			got := reflect.New(reflect.TypeOf(tt.want).Elem()).Interface()
			err = json.Unmarshal(w.Body.Bytes(), got)
			if err != nil {
				t.Fatalf("Json UnMarshal fail for: %v", w.Body.Bytes())

			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
