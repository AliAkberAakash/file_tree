package config

import (
	"encoding/json"
	"log"
	"testing"
)

func TestModelUnmarshall(t *testing.T) {
	var jsonString = []byte(`{
        "name": "Feature",  
        "type": "FD",
        "children": []
    }`)

	var fileStruct FileStruct

	err := json.Unmarshal(jsonString, &fileStruct)

	if err != nil {
		log.Println(err.Error())
		t.Errorf("Failed to unmarshal json")
	} else {
		log.Println(fileStruct)
	}
}

func TestNestedModelUnmarshall(t *testing.T) {
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

	var fileStruct FileStruct

	err := json.Unmarshal(jsonString, &fileStruct)

	if err != nil {
		log.Println(err.Error())
		t.Errorf("Failed to unmarshal json")
	} else {
		log.Println(fileStruct)
	}
}
