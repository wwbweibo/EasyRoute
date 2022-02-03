package error

import (
	"errors"
	"testing"
)

func TestNewServerError(t *testing.T) {
	err := NewServerError(200, errors.New("test"))
	if err.GetError().Error() != "test" || err.Code() != 200 {
		t.Error("NewServerError error")
	}
}

func TestFromError(t *testing.T) {
	err := FromError(errors.New("test"))
	if err.Code() != 500 {
		t.Fatal("err code should be 500")
	}
	if err.err.Error() != "test" {
		t.Fatal("err message should be test")
	}

	err = NewServerError(200, errors.New("test"))
	err2 := FromError(err)
	if err2.Code() != 200 {
		t.Fatal("err code should be 200")
	}
	if err2.err.Error() != "test" {
		t.Fatal("err message should be test")
	}
	if err2.Error() != "code: 200 err: test" {
		t.Fatal("err message should be test")
	}
}
