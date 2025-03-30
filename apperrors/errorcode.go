package apperrors

type ErrCode string

const (
	Unknown ErrCode = "U0000"

	InsertDataFailed ErrCode = "S0001"
)

func (code ErrCode) Wrap(err error, message string) *MyAppError {
	return &MyAppError{
		ErrCode: code,
		Message: message,
		Err:     err,
	}
}
