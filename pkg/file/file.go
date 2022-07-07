package file

import (
	"os"
	"strings"
)

func Generate(name string,fileExtension string) error {

	data := name + "/data"
	repo := data + "/repo"
	model := data + "/model"
	controller := name + "/controller"
	screen := name + "/screen"
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
		
		//create file for folders other than current dir,main folder and widget folder
		shouldCreateFile := folder != name && folder != data && folder != widget
		if(shouldCreateFile){
			isModelFolder := folder == model
			err := generateFile(isModelFolder,name,folder,fileExtension)
			if err != nil {
				return err
			}
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

func getFileName(path string) string{
	var pathArray = strings.Split(path,"/")
	return pathArray[len(pathArray)-1]
}

func generateFile(isModelFolder bool,featureName string,folder string,fileExtension string) error{

	if(isModelFolder){
		fullRequestFilePath := folder+"/"+featureName+"_request"+"."+fileExtension
		fullResponseFilePath := folder+"/"+featureName+"_response"+"."+fileExtension

		requestErr := createEmptyFile(fullRequestFilePath)
		responseErr := createEmptyFile(fullResponseFilePath)

		if requestErr != nil {
			return requestErr
		}
		if responseErr != nil {
			return responseErr
		}	
	}else{
		fileName := getFileName(folder)
		fullFilePath := folder+"/"+featureName+"_"+fileName+"."+fileExtension

		err := createEmptyFile(fullFilePath)

		if err != nil {
			return err
		}
	}
	return nil
}