package vault

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	config "cli_go.com/b/config"
	"gopkg.in/yaml.v3"
)

type SecretVault struct {
	Vault  string `json:"vault"`
	User   string `json:"user"`
	Secret string `json:"secret"`
}

var secrets []string

func getSecrets() {
	println("Obtendo usuarios e vault do yaml...")
	var vault string
	app := config.ParseYaml()
	secrets = app.App.Secrets[0].Values

	appYaml, errorMarshal := yaml.Marshal(app.App.Secrets)
	if errorMarshal != nil {
		log.Fatal("Erro no Marshal do Yaml!")
	}

	if strings.Contains(string((appYaml)), "vault") {
		vault = app.App.Secrets[0].Vault
	} else {
		vault = "foo"
	}
	os.Setenv("VAULT", vault)
	println("Dados obtidos com sucesso!")
}

func PopulateSecrets() {
	getSecrets()

	env := config.LoadEnv()
	println("Setando segredos para ambiente...")

	segredos := make([]*config.Segredos, len(secrets))

	for _, key := range secrets {
		os.Setenv("USER", key)
		urlCall := os.ExpandEnv("http://localhost:3000/vault/{VAULT}/secret/{USER}")

		client := &http.Client{}
		req, err := http.NewRequest("GET", urlCall, nil)
		if err == nil {
			req.Header.Set("Authorization", env.VaultAuth)

			res, errRes := client.Do(req)
			if errRes == nil {
				var resSecret SecretVault
				resBytes, errByte := ioutil.ReadAll(res.Body)
				if errByte == nil {
					error := json.Unmarshal(resBytes, &resSecret)
					if error != nil {
						log.Fatal(error)
						log.Panic("Impossível realizar o parse da resposta para struct!")
					}
					var s config.Segredos
					s.Key = key
					s.Secret = resSecret.Secret
					segredos = append(segredos, &s)
				} else {
					log.Fatal("Erro ao ler resposta do servidor do cofre!")
				}
			} else {
				log.Fatal(fmt.Errorf("Falha na requisicao de GET para o secret! Retornou status %d", res.StatusCode))
			}
		} else {
			log.Fatal(fmt.Errorf("Falha ao iniciar novo request! >>>>\n%s", err.Error()))
		}
	}
	//TODO: verificar porque as duas primeiras entradas são nulas!
	env.Secrets = segredos
	config.DefineVariables(env)
	println("Segredos setados com sucesso!")
}
