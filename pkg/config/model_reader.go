package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type ModelReader interface {
	ConvertModel(jsonString []byte, fileStruct *FileStruct) error
	ReadJsonConfig(filePath string) ([]byte, error)
	GetFileStructureFromFile(filePath string, fileStruct *FileStruct) error
}

type modelReader struct{}

func GetModelReader() ModelReader {
	return &modelReader{}
}

func (mr *modelReader) ConvertModel(jsonString []byte, fileStruct *FileStruct) error {
	err := json.Unmarshal(jsonString, fileStruct)
	return err
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

func (mr *modelReader) GetFileStructureFromFile(filePath string, fileStruct *FileStruct) error {

	jsonString, err := mr.ReadJsonConfig(filePath)

	if err != nil {
		return err
	}

	err = mr.ConvertModel(jsonString, fileStruct)

	if err != nil {
		return err
	}

	return nil
}
