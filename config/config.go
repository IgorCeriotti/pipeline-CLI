package config

import (
	"fmt"
	"io/ioutil"
	"log"

	viper "github.com/spf13/viper"
	yaml "gopkg.in/yaml.v3"
)

type Environment struct {
	VaultAuth string `yaml:"VAULT_AUTH"`
}

//TODO: metodo para ler e arquivo e retornar o struct
var Env Environment

func setViperConfig() {
	viper.SetConfigName("env")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(fmt.Errorf("Erro ao ler arquivo de env!\n%v", err))
	}

	err = viper.Unmarshal(&Env)
	if err != nil {
		log.Fatal(fmt.Errorf("Falha ao decodificar yaml de env para struct\n%v", err))
	}
}

func DefineVariables() {
	setViperConfig()

	e, err := yaml.Marshal(&Env)
	if err != nil {
		log.Fatal(fmt.Errorf("Erro ao realizar marshal do novo valor do ambiente!\n%v", err))
	}

	err = ioutil.WriteFile("./config/env.yaml", e, 0644)
	if err != nil {
		log.Fatal(fmt.Errorf("Falha ao escrever no env.yaml!\n%v", err))
	}
}
