package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type AppStruct struct {
	App struct {
		Technology        string `yaml:"technology"`
		TechnologyVersion int    `yaml:"technologyVersion"`
		UnitTesting       struct {
			Enabled         bool   `yaml:"enabled"`
			UnitTestingTool string `yaml:"unitTestingTool"`
		} `yaml:"unitTesting"`
		Secrets []struct {
			Vault  string   `yaml:"vault,omitempty"`
			Values []string `yaml:"values"`
		} `yaml:"secrets"`
	} `yaml:"app"`
}

func ParseYaml() AppStruct {
	yamlFile, err := ioutil.ReadFile("./config/app.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var app AppStruct

	err2 := yaml.Unmarshal(yamlFile, &app)
	if err2 != nil {
		log.Fatal(err2)
	}

	return app
}
