package errno

/* 错误码
 * 第 1 位 : 服务级错误码; 比如 1 为系统级错误; 2 为普通错误, 通常是由用户非法操作引起
 * 第2 3位 : 模块级错误码, 比如 01 为用户模块; 02 为订单模块
 * 第4 5位 : 具体的错误码, 比如 01 为手机号不合法; 02 为验证码输入错误
 */
var (
	OK = &Errno{Code: 0, Message: "OK"}

	ErrItemNotFound = &Errno{Code: 404, Message: "Item not found"}

	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrDatabase         = &Errno{Code: 10002, Message: "Database error"}
	ErrGetRedisList     = &Errno{Code: 10003, Message: "Get list from Redis out of expiration time"}
	ErrRewriteRedisList = &Errno{Code: 10004, Message: "rewrite list to Redis when cancel"}
	ErrRedis            = &Errno{Code: 10005, Message: "Redis error"}

	// auth
	ErrRegister          = &Errno{Code: 10101, Message: "Error occurred while registering on auth-server"}
	ErrRemoteAccessToken = &Errno{Code: 10102, Message: "Error occurred while getting oauth access token from auth-server"}
	ErrLocalAccessToken  = &Errno{Code: 10103, Message: "Error occurred while getting oauth access token from local"}
	ErrGetUserInfo       = &Errno{Code: 10104, Message: "Error occurred while getting user info from oauth-server by access token"}
	ErrFormToken         = &Errno{Code: 10105, Message: "Error occurred while forming the token"}

	// ---------------------------------------------------------------------------

	ErrBadRequest       = &Errno{Code: 20001, Message: "Request error"}
	ErrBind             = &Errno{Code: 20002, Message: "Error occurred while binding the request body to the struct."}
	ErrQuery            = &Errno{Code: 20003, Message: "Error occurred while getting url queries."}
	ErrPathParam        = &Errno{Code: 20004, Message: "Error occurred while getting path param."}
	ErrAuthToken        = &Errno{Code: 20005, Message: "Error occurred while handling the auth token"}
	ErrPermissionDenied = &Errno{Code: 20006, Message: "Permission denied."}
	ErrUploadFile       = &Errno{Code: 20007, Message: "Failed to upload file."}
	ErrGetFile          = &Errno{Code: 20008, Message: "Error occurred while get the file."}

	// user
	ErrUserNotExisted     = &Errno{Code: 20101, Message: "User not existed"}
	ErrPasswordIncorrect  = &Errno{Code: 20102, Message: "The password was incorrect."}
	ErrPasswordRepetition = &Errno{Code: 20103, Message: "The password entered twice is inconsistent"}
	ErrUserExisted        = &Errno{Code: 20104, Message: "User has existed"}
)

type Errno struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
