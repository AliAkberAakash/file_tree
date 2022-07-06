package args

import (
	"fmt"
	"os"

	"github.com/AliAkberAakash/file_tree/pkg/file"
)

func ReadArgsAndValidate() {
	args := os.Args

	fileName,fileExtension, err := validateArgs(args)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = file.Generate(fileName,fileExtension)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func validateArgs(args []string) (string,string, error) {

	var fileName string
	var fileExtension string

	if len(args) < 2 {
		return fileName,fileExtension, fmt.Errorf("Feature name is required")
	}

	fileName = args[1]
	fileExtension=args[2]

	return fileName,fileExtension, nil
}
