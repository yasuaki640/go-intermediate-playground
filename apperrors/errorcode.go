package apperrors

type ErrCode string

const (
	Unknown ErrCode = "U0000"

	InsertDataFailed ErrCode = "S0001"
	GetDataFailed    ErrCode = "S0002"
	NAData           ErrCode = "S0003"
	NoTargetData     ErrCode = "S0004"
	UpdateDataFailed ErrCode = "S0005"

	ReqBodyDecodeFailed ErrCode = "R0001"
	BadParam            ErrCode = "R0002"
)

func (code ErrCode) Wrap(err error, message string) *MyAppError {
	return &MyAppError{
		ErrCode: code,
		Message: message,
		Err:     err,
	}
}
