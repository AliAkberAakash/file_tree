package args

import (
	"fmt"
	"os"

	"github.com/AliAkberAakash/file_tree/pkg/file"
)

func ReadArgsAndValidate() {
	args := os.Args

	fileName, err := validateArgs(args)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	isFileExists,err:=checkDirectoryExists(fileName)
	if(isFileExists){	
		fmt.Println("file already exists")
	}else{
		err = file.Generate(fileName)
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
	if(err==nil){
		return true,nil
	}
	if(os.IsNotExist(err)){
		return false,nil
	}
	return false,err
}