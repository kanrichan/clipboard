package apis

const (
	ErrMissingParameterCode = iota + 10001
	ErrInvalidParameterCode
	ErrInternalErrorCode
	ErrUnAuthorizedCode
	ErrInvalidUsernameCode
	ErrInvalidUserIDCode
	ErrInvalidPasswordCode
	ErrInvalidPostIDCode
)

const (
	ErrMissingParameter = "参数缺失"
	ErrInvalidParameter = "参数错误"
	ErrInternalError    = "内部错误"
	ErrUnAuthorized     = "服务器无法验证该请求"
	ErrInvalidUsername  = "用户名已存在"
	ErrInvalidUserID    = "用户不存在"
	ErrInvalidPassword  = "用户名或密码错误"
	ErrInvalidPostID    = "内容不存在"
)
