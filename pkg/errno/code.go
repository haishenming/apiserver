package errno

var (
	// 通用错误
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server errno"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	// 用户错误
	ErrUserNotFound = &Errno{Code: 20102, Message: "The user was not found."}
)
