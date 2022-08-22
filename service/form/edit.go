package form

import (
	"github.com/MuXiFresh-be/model/form"
	"github.com/MuXiFresh-be/pkg/errno"
)

func EditForm(email string, group string, reason string,
	understand string, selfintroduction string, ifotherorganization string) error {
	if err := form.Edit(email,
		group, reason, understand,
		selfintroduction, ifotherorganization); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}
