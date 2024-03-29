package form

import (
	"github.com/MuXiFresh-be/model/form"
	"github.com/MuXiFresh-be/pkg/errno"
)

func ViewForm(email string) (*form.FormModel, error) {
	myform, err := form.View(email)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return myform, nil
}
func ViewForm1(email string) (form.FormModel, error) {
	myform, err := form.View1(email)
	if err != nil {
		return form.FormModel{}, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return myform, nil
}
