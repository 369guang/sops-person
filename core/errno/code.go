package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Msg: "OK"}
	InternalServerError = &Errno{Code: 1001, Msg: "Internal server error."}
	ErrBind             = &Errno{Code: 1002, Msg: "Error occurred while binding the request body to the struct"}
	ErrValidation       = &Errno{Code: 1003, Msg: "Validation failed."}
	ErrDatabase         = &Errno{Code: 1004, Msg: "Database error."}
	ErrToken            = &Errno{Code: 1005, Msg: "Error occurred while signing the JSON web token"}
	ErrRecordDuplicate  = &Errno{Code: 1006, Msg: "记录已存在"}
	ErrRecordNotExist   = &Errno{Code: 1007, Msg: "记录不存在"}
	ErrEncodeError      = &Errno{Code: 1008, Msg: "解码失败"}
	ErrWriteFileError   = &Errno{Code: 1009, Msg: "写入文件失败"}
	ErrYAMLEncodeError  = &Errno{Code: 1010, Msg: "YAML文件解码失败"}
	ErrMaxPage          = &Errno{Code: 1011, Msg: "超出最大分页"}
	ErrMinPage          = &Errno{Code: 1011, Msg: "最小分页不能小于1或最小分页条目不能小于1"}

	// login error
	ErrEncrypt              = &Errno{Code: 2001, Msg: "加密用户密码时错误."}
	ErrUserNotFound         = &Errno{Code: 2002, Msg: "该用户不存在."}
	ErrTokenInvalid         = &Errno{Code: 2003, Msg: "Token错误."}
	ErrPasswordIncorrect    = &Errno{Code: 2004, Msg: "密码错误."}
	ErrUserDuplicate        = &Errno{Code: 2005, Msg: "该用户已存在."}
	ErrUserOrPasswordIsNull = &Errno{Code: 2006, Msg: "用户名或密码为空."}
	ErrUserDisabled         = &Errno{Code: 2007, Msg: "用户被禁用."}
)
