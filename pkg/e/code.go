package e

type ErrCode int

// 定义错误码
//go:generate stringer -type ErrCode -linecomment -output code_string.go
const (
	SUCCESS        ErrCode = 0 // 操作成功
	ERROR          ErrCode = 1 // 操作失败
	INVALID_PARAMS ErrCode = 2 // 请求参数错误

	ERROR_WORONG_PWD     ErrCode = 10001 // 登录密码错误
	ERROR_EXIST_USERNAME ErrCode = 10002 // 用户名已存在
	ERROR_INSERT_USER    ErrCode = 10003 // 用户插入失败

	ERROR_AUTH_NOT_EXIST_SESSION ErrCode = 20001 // Session不存在
	ERROR_AUTH_INVALID_SESSION   ErrCode = 20002 // Session失效
	ERROR_AUTH_SAVE_SESSION      ErrCode = 20003 // Session保存失败
)

func (this ErrCode) Int() int {
	return int(this)
}
