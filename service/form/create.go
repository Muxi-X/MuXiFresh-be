package form

import (
	"github.com/MuXiFresh-be/model/form"
	"github.com/MuXiFresh-be/pkg/errno"
)

func CreateForm(email string, 
	group string, reason string, understand string, selfintroduction string, ifotherorganization string) (*form.FormModel, error) {
	Form, err := form.Create(email, group, reason,
		understand, selfintroduction, ifotherorganization)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return Form, nil
}
