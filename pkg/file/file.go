package file

import (
	"os"

	"github.com/AliAkberAakash/file_tree/pkg/config"
)

func Generate(featureName string, fileExtension string) error {
	mr := config.GetModelReader()

	var fs config.FileStruct
	err := mr.GetFileStructureFromFile("config.json", &fs)

	if err != nil {
		return err
	}

	// generate root folder first
	err = createFolder(featureName)
	if err != nil {
		return err
	}

	if len(fs.Children) > 0 {
		for _, fsc := range fs.Children {
			err := generateFileAndFolders(featureName, featureName, fileExtension, fsc)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func generateFileAndFolders(root string, featureName string, fileExtension string, fs config.FileStruct) error {
	if fs.Type == "FL" {
		fileName := root + "/" + featureName + "_" + fs.Name + "." + fileExtension
		err := createEmptyFile(fileName)

		if err != nil {
			return err
		}
	} else if fs.Type == "FD" {
		if len(root) > 0 {
			root = root + "/"
		}
		folderName := root + fs.Name
		err := createFolder(folderName)
		if err != nil {
			return err
		}

		if len(fs.Children) > 0 {
			for _, fsc := range fs.Children {
				err := generateFileAndFolders(folderName, featureName, fileExtension, fsc)
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
