package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type ModelReader interface {
	ConvertModel(jsonString []byte) (FileStruct, error)
	ReadJsonConfig(filePath string) ([]byte, error)
}

type modelReader struct{}

func GetModelReader() ModelReader {
	return &modelReader{}
}

func (mr *modelReader) ConvertModel(jsonString []byte) (FileStruct, error) {
	var fileStruct FileStruct
	err := json.Unmarshal(jsonString, &fileStruct)

	return fileStruct, err
}

func (mr *modelReader) ReadJsonConfig(filePath string) ([]byte, error) {
	jsonFile, err := os.Open(filePath)
	defer jsonFile.Close()

	if err != nil {
		return nil, err
	}

	byteValue, err := ioutil.ReadAll(jsonFile)

	return byteValue, err
}
