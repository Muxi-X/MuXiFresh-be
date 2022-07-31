package file

import "github.com/MuXiFresh-be/model"

// 更新图像
func (avatar *Picture) Update() error {
	return model.DB.Self.Where("email = ?", avatar.Email).Update(Picture{URL: avatar.URL}).Error
}
