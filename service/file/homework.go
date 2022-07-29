package file

import File "github.com/MuXiFresh-be/model/file"

func CreateHomework(url string, email string, homeworkID uint) error {
	var homework File.Homework
	homework.URL = url
	homework.Email = email
	homework.HomeworkID = homeworkID

	return homework.Create()
}
