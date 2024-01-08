package errno

type Error struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

var (
	ErrIncorrectAPIRoute = &Error{ErrorCode: 404, Message: "The incorrect API route"}
	InternalServerError  = &Error{ErrorCode: 500, Message: "Internal server error"}

	BadRequest        = &Error{ErrorCode: 00001, Message: "Fail."}
	ErrBind           = &Error{ErrorCode: 00001, Message: "Fail."}
	ErrPathParam      = &Error{ErrorCode: 00002, Message: "Lack Param Or Param Not Satisfiable."}
	FrequentRequests1 = &Error{ErrorCode: 00003, Message: "Post over 2 requirements within 1 minute"}
	FrequentRequests2 = &Error{ErrorCode: 00004, Message: "Post over 15 requirements within 1 day"}

	TokenInvalid        = &Error{ErrorCode: 10001, Message: "Token Invalid."}
	Unauthorized        = &Error{ErrorCode: 10002, Message: "Unauthorized."}
	BlackList           = &Error{ErrorCode: 10003, Message: "You are in our blacklist."}
	ErrAuthToken        = &Error{ErrorCode: 20005, Message: "Error occurred while handling the auth token"}
	ErrPermissionDenied = &Error{ErrorCode: 20006, Message: "Permission denied."}

	ErrPasswordIncorrect = &Error{ErrorCode: 20102, Message: "The password was incorrect."}
	ErrUserNotFound      = &Error{ErrorCode: 20103, Message: "user not found"}
)

func (err *Error) Error() string {
	return err.Message
}
