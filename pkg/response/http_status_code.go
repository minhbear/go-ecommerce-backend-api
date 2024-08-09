package response

const (
	ErrCodeSuccess      = 20001
	ErrCodeParamInvalid = 20003
	ErrInvalidToken     = 30001
	ErrCodeUserHasExit  = 40001
)

// message
var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Email is invalid",
	ErrInvalidToken:     "Invalid token",
	ErrCodeUserHasExit:  "User has exit",
}
