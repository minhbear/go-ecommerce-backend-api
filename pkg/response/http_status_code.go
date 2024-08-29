package response

const (
	ErrCodeSuccess      = 20001
	ErrCodeParamInvalid = 20003
	ErrInvalidToken     = 30001
	ErrCodeUserHasExit  = 40001
	ErrInvalidOTP       = 3002
	ErrSendMail         = 3003
	ErrInvalidParams    = 3004
)

// message
var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Email is invalid",
	ErrInvalidToken:     "Invalid token",
	ErrCodeUserHasExit:  "User has exit",
	ErrInvalidOTP:       "Invalid otp",
	ErrSendMail:         "ErrSendMail",
	ErrInvalidParams:    "Invalid params",
}
