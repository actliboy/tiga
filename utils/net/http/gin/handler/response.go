package handler

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liov/tiga/protobuf/utils/errorcode"
)

type H map[string]interface{}

type ResData struct {
	Code    errorcode.ErrCode `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	//验证码
	Details interface{} `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
}

//先信息后数据最后状态码
//入参1. data interface{},msg string,code int
//2.msg,code |data默认nil
//3.data,msg |code默认SUCCESS
//4.msg |data默认nil code默认ERROR
//5.data |msg默认"",code默认SUCCESS
func Response(ctx *gin.Context, res ...interface{}) {

	var resData ResData

	if len(res) == 1 {
		resData.Code = errorcode.Unknown
		if msgTmp, ok := res[0].(string); ok {
			resData.Message = msgTmp
			resData.Details = nil
		} else {
			resData.Details = res[0]
			resData.Code = errorcode.SUCCESS
		}
	} else if len(res) == 2 {
		if msgTmp, ok := res[0].(string); ok {
			resData.Details = nil
			resData.Message = msgTmp
			resData.Code = res[1].(errorcode.ErrCode)
		} else {
			resData.Details = res[0]
			resData.Message = res[1].(string)
			resData.Code = errorcode.SUCCESS
		}
	} else {
		resData.Details = res[0]
		resData.Message = res[1].(string)
		resData.Code = res[2].(errorcode.ErrCode)
	}

	ctx.JSON(http.StatusOK, &resData)
}

func Res(c *gin.Context, code errorcode.ErrCode, msg string, data interface{}) {
	var resData = ResData{
		Code:    code,
		Message: msg,
		Details: data,
	}
	c.JSON(http.StatusOK, &resData)
}

type File struct {
	File http.File
	Name string
}

type HttpFile interface {
	io.Reader
	Name() string
}
