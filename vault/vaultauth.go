package vault

import (
	"fmt"
	"io/ioutil"
	"log"

	"encoding/json"
	"net/http"

	conf "cli_go.com/b/config"
)

//message é retorno da api se houve erro
type Auth struct {
	Token   string `json:"token,omitempty"`
	Message string `json:"message,omitempty"`
}

func GetAuthTokenVault(user string, password string) {
	println("Obtendo auth para o cofre...")
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:3000/auth", nil)
	if err == nil {
		req.Header.Add("password", user)
		req.Header.Add("user", password)

		res, errRes := client.Do(req)

		if errRes == nil {
			var resToken Auth

			resBytes, errByte := ioutil.ReadAll(res.Body)
			if errByte == nil {
				json.Unmarshal(resBytes, &resToken)

				if len(resToken.Message) == 0 {
					conf.Env.VaultAuth = resToken.Token
					conf.DefineVariables()
				} else {
					log.Fatalf("Autenticacao no vault falhou!\n%s", resToken.Message)
				}
			} else {
				log.Fatal(fmt.Errorf("Erro ao decoficar resposta para bytes.\n%v", errByte))
			}
		} else {
			log.Panic(fmt.Errorf("Erro ao executar chamada de autenticacao no cofre!\n%v", errRes))
		}
	} else {
		log.Fatal(fmt.Errorf("Erro ao setar novo request!\n%v", err))
	}
	fmt.Println("Token de autenticacao do cofre obtido com sucesso!")
}
