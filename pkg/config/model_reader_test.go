package config

import (
	"log"
	"testing"
)

func TestReadJsonConfig(t *testing.T) {

	fileName := "/Users/bs274/go_projects/file_tree/config.json"

	log.Println("Trying to read: " + fileName)

	var modelReader = GetModelReader()

	jsonString, err := modelReader.ReadJsonConfig(fileName)

	if err != nil {
		log.Println(err.Error())
		t.Error("Failed to read config file")
	} else {
		log.Println(string(jsonString))
	}

}
