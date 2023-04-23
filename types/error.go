package types

const (
	GptFailedError = 1001
	UserLoginError = 1002
)

// GptError is a custom error type
// It contains the error message and the error code
type GptError struct {
	Err  error
	Code int
}

func (e *GptError) Error() string {
	return e.Err.Error()
}

// LoginError is a custom error type
// It contains the error message and the error code
// This error need user login first
type LoginError struct {
	Err  error
	Code int
}

func (e *LoginError) Error() string {
	return e.Err.Error()
}
