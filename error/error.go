package error

import "strconv"

type ServerError struct {
	code int
	err  error
}

func NewServerError(code int, err error) *ServerError {
	return &ServerError{
		code: code,
		err:  err,
	}
}

func (err *ServerError) Code() int {
	return err.code
}

func (err *ServerError) GetError() error {
	return err.err
}

func (err *ServerError) Error() string {
	return "code: " + strconv.Itoa(err.code) + " err: " + err.err.Error()
}

func FromError(err error) *ServerError {
	if e, ok := err.(*ServerError); ok {
		return e
	} else {
		return NewServerError(500, err)
	}
}
