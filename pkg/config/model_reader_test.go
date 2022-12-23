package config

import (
	"log"
	"testing"
)

func TestReadJsonConfig(t *testing.T) {

	fileName := "/Users/bs274/go_projects/file_tree/config.json"

	var modelReader = GetModelReader()

	jsonString, err := modelReader.ReadJsonConfig(fileName)

	if err != nil {
		log.Println(err.Error())
		t.Error("Failed to read config file")
	} else {
		log.Println(string(jsonString))
	}

}

func TestConvertModel(t *testing.T) {

	var jsonString = []byte(`{
        "name": "Feature",  
        "type": "FD",
        "children": [
			{
				"name": "Feature",  
				"type": "FD",
				"children": []
			},
			{
				"name": "Feature",  
				"type": "FD",
				"children": [
					{
						"name": "Feature",  
						"type": "FD",
						"children": []
					}
				]
			}
		]
    }`)

	var modelReader = GetModelReader()

	output, err := modelReader.ConvertModel(jsonString)

	if err != nil {
		log.Println(err.Error())
		t.Error("Failed to convert model")
	} else {
		log.Println(output)
	}

}

func TestModelReaderComplete(t *testing.T) {
	fileName := "/Users/bs274/go_projects/file_tree/config.json"

	var modelReader = GetModelReader()

	jsonString, err := modelReader.ReadJsonConfig(fileName)

	if err != nil {
		log.Println(err.Error())
		t.Error("Failed to read config file")
	} else {
		log.Println(string(jsonString))
	}

	output, err := modelReader.ConvertModel(jsonString)

	if err != nil {
		log.Println(err.Error())
		t.Error("Failed to convert model")
	} else {
		log.Println(output)
	}
}
