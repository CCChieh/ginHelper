package ginHelper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var GenHandlerFunc gin.HandlerFunc = nil

type Data interface{}

type Parameter interface {
	Bind(c *gin.Context, p Parameter) (err error)  //绑定参数
	Handler(c *gin.Context) (data Data, err error) //执行具体业务
	Result(c *gin.Context, data Data, err error)   //结果返回
}

type BaseParam struct {
}

func (param *BaseParam) Bind(c *gin.Context, p Parameter) (err error) {
	if err := c.ShouldBind(p); err != nil {
		return err
	}
	if err := c.ShouldBindUri(p); err != nil {
		return err
	}
	if err := c.ShouldBindHeader(p); err != nil {
		return err
	}
	return err
}

func (param *BaseParam) Handler(c *gin.Context) (data Data, err error) {
	return param, nil
}

func (param *BaseParam) Result(c *gin.Context, data Data, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, data)
	}
}
