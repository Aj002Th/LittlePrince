package e

// 内部错误码
var MsgFlags = map[int]string{
	SUCCESS:        "成功",
	ERROR:          "失败",
	INVALID_PARAMS: "请求参数错误",
	UNAUTHORIZED:   "未登录",

	ERROR_AUTH:      "未进行身份验证",
	ERROR_FORBIDDEN: "身份验证不通过",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
