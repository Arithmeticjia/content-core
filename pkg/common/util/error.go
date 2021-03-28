package util

var BadContentType = NewError("D", "01", "01")

type ContentError struct {
	module string
	operation string
	code string
}

func NewError(module, operation, code string) ContentError {
	return ContentError{
		module,
		operation,
		code,
	}
}

func (e ContentError) Error() string {
	return e.String()
}

func (e ContentError) String() string {
	return e.module + e.operation + e.code
}