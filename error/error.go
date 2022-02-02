package error

import "strconv"

type ServerError struct {
	code int
	err  string
}

func NewServerError(code int, err string) *ServerError {
	return &ServerError{
		code: code,
		err:  err,
	}
}

func (err *ServerError) GetCode() int {
	return err.code
}

func (err *ServerError) GetError() string {
	return err.err
}

func (err *ServerError) Error() string {
	return "code: " + strconv.Itoa(err.code) + " err: " + err.err
}

func FromError(err error) *ServerError {
	if e, ok := err.(*ServerError); ok {
		return e
	} else {
		return NewServerError(500, err.Error())
	}
}
