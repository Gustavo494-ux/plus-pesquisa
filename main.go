package main

import (
	"fmt"
	"strconv"

	"plus-pesquisa/api"
	"plus-pesquisa/api/model"
	"plus-pesquisa/api/utils"
)

func main() {
	GerarDefinirHash() //Preenche o hash do id de todas as pesquisas.

	api.ConfigurarApi()
}

/*
Pode ignorar essa função. acredito que você não irá precisar fazer isso no projeto. irá precisar gerar o hash e salvar apenas no banco de dados acredito.
Quando for realiza o cadastro na api a gente conversa sobre isso.
*/
func GerarDefinirHash() {
	var err error
	for i := range model.Pesquisas {
		model.Pesquisas[i].Id_Hash, err = utils.GerarHash(strconv.Itoa(int(model.Pesquisas[i].Id)))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(model.Pesquisas[i])
	}
}
