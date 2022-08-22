package form

import (
	"github.com/MuXiFresh-be/model/form"
)

func SearchForm(group string) ([]form.FormModel, error) {
	var forms []form.FormModel
	forms, err := form.Search(group)
	if err != nil {
		return forms, err
	}
	return forms, nil
}
