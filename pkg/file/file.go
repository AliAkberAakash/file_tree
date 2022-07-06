package file

import (
	"os"
)

func Generate(name string,fileExtension string) error {

	data := name + "/data"
	repo := data + "/repo"
	model := data + "/model"
	controller := name + "/controller"
	screen := name + "/screeen"
	widget := screen + "/widget"

	var folders []string

	folders = append(folders, name)
	folders = append(folders, data)
	folders = append(folders, repo)
	folders = append(folders, model)
	folders = append(folders, controller)
	folders = append(folders, screen)
	folders = append(folders, widget)

	for _, folder := range folders {
		err := createFolder(folder)
		if err != nil {
			return err
		}	
	}

	return nil
}

func createEmptyFile(name string) error {
	d := []byte("")
	return os.WriteFile(name, d, 0644)
}

func createFolder(name string) error {
	return os.Mkdir(name, os.ModePerm)
}
