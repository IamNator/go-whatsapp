package go_whatsapp

type MetaError string

func (e MetaError) Error() string {
	return string(e)
}

const (
	ErrAPICallLimitExceeded MetaError = "call limit exceeded"
)

var (
	errCodeMap = map[MetaError]int{
		ErrAPICallLimitExceeded: 80007,
	}

	codeErrorMap = map[int]MetaError{
		80007: ErrAPICallLimitExceeded,
	}
)

func (e MetaError) Code() int {
	return errCodeMap[e]
}

func Error(code int) MetaError {
	return codeErrorMap[code]
}

func IsError(err error) bool {
	_, ok := err.(MetaError)
	return ok
}
