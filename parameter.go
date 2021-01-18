package ginHelper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var GenHandlerFunc gin.HandlerFunc = nil

type Parameter interface {
	Error() error         //错误返回
	Mount(c *gin.Context) //挂载Context
	BeforeBind()          //绑定参数前的操作
	Bind()                //绑定参数
	Service()             //执行具体业务
	Result()              //结果返回
}

type Param struct {
	Context *gin.Context //上下文
	Err     error        //存储内部产生的错误
	Ret     interface{}  //存储返回的结构体
}

func (param *Param) Mount(c *gin.Context) {
	param.Context = c
}

func (param *Param) BeforeBind() {
}

func (param *Param) Error() error {
	return param.Err
}

func (param *Param) Bind() {
	param.Err = param.Context.ShouldBind(param)
}

func (param *Param) Service() {
}

func (param *Param) Result() {
	if param.Err != nil {
		param.Context.JSON(http.StatusBadRequest, gin.H{"message": param.Err.Error()})
	} else {
		param.Context.JSON(http.StatusOK, param.Ret)
	}
}
