package ginHelper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var GenHandlerFunc gin.HandlerFunc = nil

// Context 似乎只能通过这种方式传输进来
type Parameter interface {
	Error() error                     //错误返回
	BeforeBind(c *gin.Context)        //绑定参数前的操作
	Bind(c *gin.Context, p Parameter) //绑定参数
	AfterBind(c *gin.Context)         //绑定参数后操作
	Service()                         //执行具体业务
	Result(c *gin.Context)            //结果返回
}

type Param struct {
	Err error       //存储内部产生的错误
	Ret interface{} //存储返回的结构体
}

func (param *Param) BeforeBind(c *gin.Context) {
}

func (param *Param) AfterBind(c *gin.Context) {
}

func (param *Param) Error() error {
	return param.Err
}

func (param *Param) Bind(c *gin.Context, p Parameter) {
	param.Err = c.ShouldBind(p)
}

func (param *Param) Service() {
}

func (param *Param) Result(c *gin.Context) {
	if param.Err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": param.Err.Error()})
	} else {
		c.JSON(http.StatusOK, param.Ret)
	}
}
