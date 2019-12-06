# ginHelper

`ginHelper`是一个自动把`handlerFunc`添加到`Gin`并且设置路由的工具包。

1. 在存放`handlerFunc`的文件夹中任意位置写下
```go
type helper struct {
}
```  
2. 之后每次写`handlerFunc`的时候都类似下方的`helloHandler`前面加上一个
`*helper`的一个方法`HelloHandler()`中设置路由。

```go
func (h *Helper) HelloHandler() (r *ginHelper.Router) {
    return &ginHelper.Router{
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
3. 在运行gin的时候

```go
r := gin.New()
```
自动导入handler文件夹中的所有路由
```go
ginHelper.Build(new(handler.Helper), r)
```
将user文件夹中所有的路由导入到同一个路由组中
```go
ginHelper.Build(new(user.Helper), r.Group("/user"))
```