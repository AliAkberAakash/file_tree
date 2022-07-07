package args

import (
	"fmt"
	"os"

	"github.com/AliAkberAakash/file_tree/pkg/file"
	"github.com/manifoldco/promptui"
)

func ReadArgsAndValidate() {
	args := os.Args

	fileName, fileExtension, err := validateArgs(args)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = file.Generate(fileName, fileExtension)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func validateArgs(args []string) (string, string, error) {

	var fileName string
	var fileExtension string

	if len(args) < 2 {
		return fileName, fileExtension, fmt.Errorf("Feature name is required")
	}

	fileName = args[1]
	fileExtension = args[2]

	return fileName, fileExtension, nil
}

func checkDirectoryExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func displayPrompt(fileName string) (bool, error) {
	var overWriteDirectory bool = false
	//configure prompt
	prompt := promptui.Select{
		Label: "Do you want to replace the contents in " + fileName + " directory",
		Items: []string{"Yes", "No"},
	}

	//display prompt
	_, result, err := prompt.Run()

	//if the user selects Yes
	if result == "Yes" {
		overWriteDirectory = true
	}

	return overWriteDirectory, err

}
