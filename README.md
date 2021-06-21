# ginHelper

之前接触过Java的Swagger，非常简单易用，但是Golang中一直没有合适的Swagger实现。现在比较流行的方案是使用注释实现，需要手动写注释，总感觉不是特别方便，所以就自己实现了个该工具。坑是挖了很久，刚学go的时候挖的，之后无奈搬砖去写java，一直没机会填坑，最近又重新写go了，所以回来慢慢填坑。

`ginHelper`支持的功能：

* 整合gin的参数绑定与路由设置
* 非注释自动生成swagger


## 路由与参数

为了自动绑定参数和生成路由，需要先了解下面几个概念：

* GroupRouter 路由组

```go
type GroupRouter struct {
	Path   string   // 路由组的根路径，与Gin的Group一样，定义一组接口的公共路径
	Name   string   // 路由组的名称
	Routes []*Route // 路由组中的具体路由
}
```

定义一组相关的路由

* Router 路由

```go
type Route struct {
	Param    Parameter         // 接口的参数实现
	Path     string            // 接口的路径
	Method   string            // 接口的方法
	Handlers []gin.HandlerFunc // 接口的额外处理函数
}
```

* 参数绑定

为了成功绑定参数，并降低代码的重复度，需要参数实现`Parameter`接口：

```go
type Parameter interface {
	Bind(c *gin.Context, p Parameter) (err error)  //绑定参数
	Handler(c *gin.Context) (data Data, err error) //执行具体业务
	Result(c *gin.Context, data Data, err error)   //结果返回
}
```

为了避免不必要的重复实现接口，可以使用结构体嵌入，比如使用内置的`BaseParam`嵌入。

## 基本使用

包内包含两种初始化方法：不需要Swagger的`func New() *Helper`,和自动生成swagger的`func NewWithSwagger(swaggerInfo *SwaggerInfo, r GinRouter) *Helper`

示例：

```go
// 定义一个Group
var testGroup = &ginHelper.GroupRouter{
	Path: "test",
	Name: "Mytest",
	Routes: []*ginHelper.Route{
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

// 接口的参数
type testBodyParam struct {
	ginHelper.BaseParam `json:"-"`
	Foo       string    `binding:"required"`
	FooName   string    `json:"fName" binding:"required"`
	FooInt    int       `binding:"required"`
	FooIgnore string    `json:"-"`
	FooStruct
	FooStruct2 FooStruct
	FooStruct3 *FooStruct
}

func Example() {
	router := gin.Default()
	r := router.Group("api")
    // 如果不需要swagger，可以使用New初始化
	h := ginHelper.NewWithSwagger(&ginHelper.SwaggerInfo{
		Description: "swagger test page",
		Title:       "Swagger Test Page",
		Version:     "0.0.1",
		ContactInfoProps: ginHelper.ContactInfoProps{
			Name:  "zzj",
			URL:   "https://zzj.cool",
			Email: "email@zzj.cool",
		},
	}, r)
	h.Add(testGroup, r)
	_ = router.Run(":8888")
}
```

如果开启了swagger的话，访问`http://127.0.0.1:8888/api/swagger`即可。

性能测试：

直接使用Gin和使用GinHelper生成的接口，两者性能差别不大：

go test -bench=. -benchmem -run=none

```shell
goos: linux
goarch: amd64
pkg: github.com/ccchieh/ginHelper
cpu: AMD Ryzen 5 3400G with Radeon Vega Graphics
BenchmarkHelp-8           236252              4836 ns/op            2258 B/op         38 allocs/op
BenchmarkNorm-8           254684              4765 ns/op            2258 B/op         38 allocs/op
PASS
ok      github.com/ccchieh/ginHelper    2.466s
```
