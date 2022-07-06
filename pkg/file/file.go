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
	var fileName string

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
		if(folder != name && folder != data && folder != widget ){
			if(folder == model){
				fullRequestFilePath:=folder+"/"+name+"_request"+"."+fileExtension
				fullResponseFilePath2:=folder+"/"+name+"_response"+"."+fileExtension
				requestErr:= createEmptyFile(fullRequestFilePath)
				responseErr:= createEmptyFile(fullResponseFilePath2)
				if requestErr != nil {
					return err
				}
				if responseErr != nil {
					return err
				}	
			}else{
				fileName=getFileName(folder)
				fullFilePath:=folder+"/"+name+"_"+fileName+"."+fileExtension
				err:= createEmptyFile(fullFilePath)
				if err != nil {
					return err
				}
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
