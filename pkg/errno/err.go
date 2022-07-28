package errno

import (
	errors "github.com/micro/go-micro/errors"
)

func ServerErr(errno *Errno, cause string) error {
	return &errors.Error{
		// Id:     "",
		Code:   int32(errno.Code),
		Detail: cause,
		Status: errno.Message,
	}
}

func NotFoundErr(errno *Errno, cause string) error {
	return &errors.Error{
		Code:   404,
		Detail: cause,
		Status: errno.Message,
	}
}

func (e Errno) Error() string {
	return e.Message
}

func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Errno:
		return typed.Code, typed.Message
	case *errors.Error:
		return int(typed.Code), typed.Status + " : " + typed.Detail
	default:
		return InternalServerError.Code, err.Error()
	}
}
