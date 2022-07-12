package response

import (
	"net/http"

	"github.com/Aj002Th/LittlePrince/pkg/e"
	"github.com/Aj002Th/LittlePrince/pkg/logging"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// 统一的响应格式
func Result(code int, data interface{}, msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(ctx *gin.Context) {
	Result(e.SUCCESS.Int(), map[string]interface{}{}, e.SUCCESS.String(), ctx)
}

func OkWithMessage(message string, ctx *gin.Context) {
	Result(e.SUCCESS.Int(), map[string]interface{}{}, message, ctx)
}

func OkWithData(data interface{}, ctx *gin.Context) {
	Result(e.SUCCESS.Int(), data, e.SUCCESS.String(), ctx)
}

func OkWithDetailed(data interface{}, message string, ctx *gin.Context) {
	Result(e.SUCCESS.Int(), data, message, ctx)
}

func Fail(code e.ErrCode, ctx *gin.Context) {
	logging.Error(code.String())
	Result(code.Int(), map[string]interface{}{}, code.String(), ctx)
}

func FailWithMessage(code e.ErrCode, message string, ctx *gin.Context) {
	Result(code.Int(), map[string]interface{}{}, message, ctx)
}

func FailWithData(code e.ErrCode, data interface{}, ctx *gin.Context) {
	Result(code.Int(), data, code.String(), ctx)
}

func FailWithDetailed(code e.ErrCode, data interface{}, message string, ctx *gin.Context) {
	Result(code.Int(), data, message, ctx)
}
