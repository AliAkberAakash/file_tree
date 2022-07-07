package args

import (
	"fmt"
	"log"
	"os"

	"github.com/AliAkberAakash/file_tree/pkg/file"
	"github.com/manifoldco/promptui"
)

func ReadArgsAndValidate() {
	args := os.Args

	fileName, err := validateArgs(args)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//check if the directory already exists
	isFileExists,err := checkDirectoryExists(fileName)
	if(isFileExists){	
		shouldOverwrite,err := displayPrompt(fileName)

		if err != nil {
			log.Fatalf("Prompt failed %v\n", err)		
		}

		if(shouldOverwrite){
			os.RemoveAll(fileName)
			generateFile(fileName)
		}
		
	}else{
		generateFile(fileName)
	}

	if err != nil {
		fmt.Println(err.Error())
	}
}

func validateArgs(args []string) (string, error) {

	var fileName string

	if len(args) < 2 {
		return fileName, fmt.Errorf("Feature name is required")
	}

	fileName = args[1]

	return fileName, nil
}

func checkDirectoryExists(path string)(bool,error){
	_,err:=os.Stat(path)
	if err==nil {
		return true,nil
	}
	if os.IsNotExist(err) {
		return false,nil
	}
	return false,err
}

func generateFile(fileName string){
	err:=file.Generate(fileName)
	if err != nil{
		fmt.Println(err.Error())
	}
}

func displayPrompt(fileName string) (bool,error){
	var overWriteDirectory bool = false
	//configure prompt
	prompt := promptui.Select{
		Label: "Do you want to replace the contents in "+fileName+" directory",
		Items: []string{"Yes", "No"},
	}

	//display prompt
	_, result, err := prompt.Run()
	
	//if the user selects Yes
	if result == "Yes" {
		overWriteDirectory = true
	}

	return overWriteDirectory,err

}