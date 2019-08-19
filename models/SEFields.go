package models

import (
	"io/ioutil"
	"log"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

//Sefields is the searchable and uneditable fields
var Sefields *SEFields

//SEFields is a struct that grabs the searchable and editable fields
type SEFields struct {
	Searchable []string
	Uneditable []string
}

//GetSEFields gets two arrays
func (g *SEFields) GetSEFields() *SEFields {
	path, _ := filepath.Abs("./models/SEFields.yaml")
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, g)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return g
}
