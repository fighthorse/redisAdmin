package self_errors

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ErrCode struct {
	Code int
	Msg  string
}

var (
	JsonErr = ErrCode{Code: 10000, Msg: "参数解析失败"}

	ParamsErr = ErrCode{Code: 10000, Msg: "参数校验失败"}
)

func JsonErrExport(data ErrCode, err error, userMsg string) map[string]interface{} {
	errMsg := ""
	if err != nil {
		errMsg = ":" + err.Error()
	}
	if userMsg != "" {
		return gin.H{"code": data.Code, "message": fmt.Sprintf("%s%s", userMsg, errMsg), "data": map[string]interface{}{}}
	}
	return gin.H{"code": data.Code, "message": fmt.Sprintf("%s%s", data.Msg, errMsg), "data": map[string]interface{}{}}

}
