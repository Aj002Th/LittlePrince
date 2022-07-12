// Code generated by "stringer -type ErrCode -linecomment -output code_string.go"; DO NOT EDIT.

package e

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SUCCESS-0]
	_ = x[ERROR-1]
	_ = x[INVALID_PARAMS-2]
	_ = x[ERROR_WORONG_PWD-10001]
	_ = x[ERROR_EXIST_USERNAME-10002]
	_ = x[ERROR_INSERT_USER-10003]
	_ = x[ERROR_AUTH_NOT_EXIST_SESSION-20001]
	_ = x[ERROR_AUTH_INVALID_SESSION-20002]
	_ = x[ERROR_AUTH_SAVE_SESSION-20003]
}

const (
	_ErrCode_name_0 = "操作成功操作失败请求参数错误"
	_ErrCode_name_1 = "登录密码错误用户名已存在用户插入失败"
	_ErrCode_name_2 = "Session不存在Session失效Session保存失败"
)

var (
	_ErrCode_index_0 = [...]uint8{0, 12, 24, 42}
	_ErrCode_index_1 = [...]uint8{0, 18, 36, 54}
	_ErrCode_index_2 = [...]uint8{0, 16, 29, 48}
)

func (i ErrCode) String() string {
	switch {
	case 0 <= i && i <= 2:
		return _ErrCode_name_0[_ErrCode_index_0[i]:_ErrCode_index_0[i+1]]
	case 10001 <= i && i <= 10003:
		i -= 10001
		return _ErrCode_name_1[_ErrCode_index_1[i]:_ErrCode_index_1[i+1]]
	case 20001 <= i && i <= 20003:
		i -= 20001
		return _ErrCode_name_2[_ErrCode_index_2[i]:_ErrCode_index_2[i+1]]
	default:
		return "ErrCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
