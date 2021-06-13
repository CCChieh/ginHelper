# ginHelper

`ginHelper`支持的功能：

* 利用反射自动检索某一包内所有的`HandlerFunc`(即gin的Handler)添加到`Gin`的路由中。
* 引入`parameter`接口的概念，通过定义实现该接口的结构体后可以自动生成一个`gin.HandlerFunc`,利用该特点可以极大地降低handler和service的代码重复。

示例在<https://github.com/CCChieh/ginHelper_example>

## 自动检索包下的handler

1. 在存放`handlerFunc`的文件夹中任意位置写下

    ```go
    type helper struct {
    }
    ```

2. 之后每次写`handlerFunc`的时候都类似下方的`helloHandler`前面加上一个
    `*helper`的一个方法`HelloHandler()`中设置路由。

    ```go
    func (h *Helper) HelloHandler() (r *ginHelper.Router) {
        handler := func(c *gin.Context) {
            c.String(http.StatusOK, "Hello world!")
        }
        return &ginHelper.Router{
            Path:   "/HelloHandler",
            Method: "GET",
            Handlers: []gin.HandlerFunc{
                handler,
            },
        }
    }

    ```

## `parameter`接口

### `parameter`接口的定义

```go
type parameter interface {
 Error() error                     //错误返回
 BeforeBind(c *gin.Context)        //绑定参数前的操作
 Bind(c *gin.Context, p parameter) //绑定参数
 AfterBind(c *gin.Context)         //绑定参数后操作
 Service(c *gin.Context)                         //执行具体业务
 Result(c *gin.Context)            //结果返回
}
```

具体执行的流程：

![](https://raw.githubusercontent.com/CCChieh/image/master/%E6%B5%81%E7%A8%8B%E5%9B%BE.png)

### `ginHelper`中内置的Param结构体

在`ginHelper`中实现了一个最基本的`parameter`接口的`Param`结构体:

```go
type Param struct {
 Err error //存储内部产生的错误
 Ret interface{} //存储返回的结构体
}

func (param *Param) BeforeBind(c *gin.Context) {
}

func (param *Param) AfterBind(c *gin.Context) {
}

func (param *Param) Error() error {
 return param.Err
}

func (param *Param) Bind(c *gin.Context, p parameter) {
 param.Err = c.ShouldBind(p)
}

func (param *Param) Service(c *gin.Context) {
}

func (param *Param) Result(c *gin.Context) {
 if param.Err != nil {
  c.JSON(http.StatusBadRequest, gin.H{"message": param.Err.Error()})
 } else {
  c.JSON(http.StatusOK, param.Ret)
 }
}
```

### 使用`parameter`接口实现handler以及service

1. 自定义一个参数结构体

   这里自定义一个Hello结构体，引用结构体`ginHelper.Param`之后 重写Service方法。

   ```go
   type Hello struct {
    ginHelper.Param
    Name      string `form:"name" binding:"required"`
   }
   
   func (param *Hello) Service() {
    ginHelper.Ret = gin.H{"message": "Hello " + param.Name + "!"}
   }
   ```

   注意这里`Name`的tag使用参考gin的[模型绑定和验证](https://gin-gonic.com/zh-cn/docs/examples/binding-and-validation/)，对于其他方法参见`parameter`接口，可以重写该接口的所有方法来实现自定义。

   当然也可以不使用`ginHelper.Param`自己实现`parameter`接口的所有方法。可以根据自己的需求选择不同的方式。

2. 进一步简化`HelloHandler()`

   在引进了`parameter`接口后`HelloHandler()`将被进一步简化为：

   ```go
   func (h *Helper) HelloHandler() (r *ginHelper.Router) {
    return &ginHeXlper.Router{
     Param:  new(service.Hello),
     Path:   "/",
     Method: "GET",
    }
   }
   ```

3. 使用中间件

   这里内置了个变量`ginHelper.GenHandlerFunc`其值为`nil`，可以用来表示将会自动生成的handler，所以在使用中间件后的形式可以为：

   ```go
   func (h *Helper) AdminHandler() (r *ginHelper.Router) {
    return &ginHelper.Router{
     Param:  new(service.Hello),
     Path:   "/admin",
     Method: "GET",
     Handlers: []gin.HandlerFunc{
      middleware.AdminMiddleware(),
      ginHelper.GenHandlerFunc,
     },
    }
   }
   ```

   `middleware.AdminMiddleware()`是一个身份验证的中间件，这样子就可以定义中间件和自动生成的handler之间的顺序关系，当然如果默认自动生成的handler放到最后的话，写法还可以省略为：

   ```go
   func (h *Helper) AdminHandler() (r *ginHelper.Router) {
    return &ginHelper.Router{
     Param:  new(service.Hello),
     Path:   "/admin",
     Method: "GET",
     Handlers: []gin.HandlerFunc{
      middleware.AdminMiddleware(),
     },
    }
   }
   ```

## 使用`ginHelper`自动构建路由

在运行gin的时候

```go
r := gin.New()
```

自动导入handler包中的所有路由

```go
ginHelper.Build(new(handler.Helper), r)
```

将user包中所有的路由导入到同一个路由组中

```go
ginHelper.Build(new(user.Helper), r.Group("/user"))
```

性能测试：

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
