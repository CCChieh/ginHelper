# ginHelper

`ginHelper`是一个自动把`handlerFunc`添加到`Gin`并且设置路由的工具包。

1. 在handler文件夹中添加`main.go`文件
```
type helper struct {
}

func Build(r gin.IRoutes) {
    h := new(helper)
    valueOfh := reflect.ValueOf(h)
    numMethod := valueOfh.NumMethod()
    for i := 0; i < numMethod; i++ {
        rt := valueOfh.Method(i).Call(nil)[0].Interface().(*handlerHelper.Router)
        rt.AddHandler(r)
    }
}
```  
2. 之后每次写`handlerFunc`的时候都类似下方的`helloHandler`前面加上一个
`*helper`的一个方法`HelloHandler()`中设置路由。

```
func (h *helper) HelloHandler() (r *handlerHelper.Router) {
    return &handlerHelper.Router{
        Path:   "/HelloHandler",
        Method: "GET",
        Handlers: []gin.HandlerFunc{
            helloHandler,
        }}
}
func helloHandler(c *gin.Context) {

    c.String(http.StatusOK, "Hello world!")
}
```
3. 在gin中调用

```go
r := gin.New()
//handler中自动建立路由
handler.Build(r)
//user的Group中自动建立路由
user.Build(r.Group("/user"))
```