package conf

import (
	"io/ioutil"
	"log"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

//Conf is an object created by the conf.yaml file
type Conf struct {
	Database Database
	Keycloak Keycloak
	App      App
}

//Database Config Object
type Database struct {
	Host       string
	Name       string
	Selectuser string
	Selectpass string
	Updateuser string
	Updatepass string
}

//Keycloak Config Object
type Keycloak struct {
	Host string
}

//App Config Object
type App struct {
	Port int
	Host string
}

//GetConf fills the conf struct
func (c *Conf) GetConf() *Conf {
	path, _ := filepath.Abs("./conf/conf.yaml")
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
