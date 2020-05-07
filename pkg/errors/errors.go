package errors

type CommandError interface {
	error
	Code() int
}

type GeneralError struct {
	Message string
}

func (GeneralError) Code() int {
	return 1
}

func (e GeneralError) Error() string {
	return e.Message
}

func ConvError(err error) CommandError {
	return GeneralError{Message: err.Error()}
}
