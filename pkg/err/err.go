package err

import (
	"encoding/json"
	"net/http"

	errors "github.com/micro/go-micro/errors"
)

type ErrDetail struct {
	Errno
	Cause string `json:"cause"`
}

func NotFoundErr(errno *Errno, cause string) error {
	detail := &ErrDetail{
		*errno,
		cause,
	}

	detailStr, _ := json.Marshal(detail)

	return &errors.Error{
		Id:     "serverName",
		Code:   404,
		Detail: string(detailStr),
		Status: http.StatusText(404),
	}
}

func BadRequestErr(errno *Errno, cause string) error {
	detail := &ErrDetail{
		*errno,
		cause,
	}

	detailStr, _ := json.Marshal(detail)

	return &errors.Error{
		Id:     "serverName",
		Code:   400,
		Detail: string(detailStr),
		Status: http.StatusText(400),
	}
}

func ServerErr(errno *Errno, cause string) error {
	detail := &ErrDetail{
		*errno,
		cause,
	}

	detailStr, _ := json.Marshal(detail)

	return &errors.Error{
		Id:     "serverName",
		Code:   500,
		Detail: string(detailStr),
		Status: http.StatusText(500),
	}
}

func ParseDetail(detail string) (*ErrDetail, error) {
	d := &ErrDetail{}
	err := json.Unmarshal([]byte(detail), d)
	if err != nil {
		return nil, err
	}
	return d, nil
}
