package pkg

type Errors struct {
	errorMessage string
	errorCode    int
}

func NewError(message string, code int) *Errors {
	return &Errors{
		errorMessage: message,
		errorCode:    code,
	}
}

func (e *Errors) Error() string {
	return e.errorMessage
}

func (e *Errors) Status() int {
	return e.errorCode
}
