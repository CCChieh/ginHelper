package ginHelper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var GenHandlerFunc gin.HandlerFunc = nil

type Data interface{}

// Context 似乎只能通过这种方式传输进来
type Parameter interface {
	Bind(c *gin.Context, p Parameter) (err error)  //绑定参数
	Service(c *gin.Context) (data Data, err error) //执行具体业务
	Result(c *gin.Context, data Data, err error)   //结果返回
}

type BaseParam struct {
}

func (param *BaseParam) Bind(c *gin.Context, p Parameter) (err error) {
	if err := c.ShouldBind(p); err != nil {
		return err
	}
	return err
}

func (param *BaseParam) Service(c *gin.Context) (data Data, err error) {
	return nil, nil
}

func (param *BaseParam) Result(c *gin.Context, data Data, err error) {
	if err != nil {
		c.String(http.StatusBadRequest, "%s", gin.H{"message": err.Error()})
	} else {
		c.String(http.StatusOK, "%s", data)
	}
}
