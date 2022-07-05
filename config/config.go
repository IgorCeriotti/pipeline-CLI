package config

import (
	"fmt"
	"io/ioutil"
	"log"

	viper "github.com/spf13/viper"
	yaml "gopkg.in/yaml.v3"
)

type Environment struct {
	VaultAuth string      `yaml:"VAULT_AUTH"`
	Secrets   []*Segredos `yaml:"secrets,omitempty"`
}

type Segredos struct {
	Key    string `yaml:"key"`
	Secret string `yaml:"secret"`
}

func setViperConfig(env Environment) {
	viper.SetConfigName("env")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(fmt.Errorf("Erro ao ler arquivo de env!\n%v", err))
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal(fmt.Errorf("Falha ao decodificar yaml de env para struct\n%v", err))
	}
}

func DefineVariables(env Environment) {
	setViperConfig(env)

	e, err := yaml.Marshal(&env)
	if err != nil {
		log.Fatal(fmt.Errorf("Erro ao realizar marshal do novo valor do ambiente!\n%v", err))
	}

	err = ioutil.WriteFile("./config/env.yaml", e, 0644)
	if err != nil {
		log.Fatal(fmt.Errorf("Falha ao escrever no env.yaml!\n%v", err))
	}
}

func LoadEnv() Environment {
	envFile, err := ioutil.ReadFile("./config/env.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var env Environment

	err2 := yaml.Unmarshal(envFile, &env)
	if err2 != nil {
		log.Fatal(err2)
	}

	return env
}
