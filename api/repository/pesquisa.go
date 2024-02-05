package repository

import "plus-pesquisa/api/model"

/*
	O Pacote repository é responsável pelas funções do banco de dados.
	Neste exemplo irei simular a existência do banco de dados com um slice(Um array de tamanho dinamico)
*/

// VerificarSePesquisaRespondida: verifica no banco se a pesquisa foi respondida
func VerificarSePesquisaRespondida(id uint64) bool {
	for _, pesquisa := range model.Pesquisas {
		if pesquisa.Id == id {
			if pesquisa.Nota > 0 {
				return true
			}
		}
	}
	return false
}

// BuscarPesquisa: busca os dados da pesquisa
func BuscarPesquisa(id uint64) model.Pesquisa {
	for _, pesquisa := range model.Pesquisas {
		if pesquisa.Id == id {
			return pesquisa
		}
	}
	return model.Pesquisa{} // retorna uma pesquisa vazia
}

// ResponderPesquisa: responde a pesquisa
func ResponderPesquisa(id, nota uint64) {
	for i, pesquisa := range model.Pesquisas {
		if pesquisa.Id == id {
			model.Pesquisas[i].Nota = uint(nota)
		}
	}
}

/*******************************************Implementação das funções acima porem com o hash do id ao invez do id***********************************************/
// VerificarSePesquisaRespondidaHash: verifica no banco se a pesquisa foi respondida utilizando o hash do id
func VerificarSePesquisaRespondidaHash(Id_Hash string) bool {
	for _, pesquisa := range model.Pesquisas {
		if pesquisa.Id_Hash == Id_Hash {
			if pesquisa.Nota > 0 {
				return true
			}
		}
	}
	return false
}

// BuscarPesquisaHash: busca os dados da pesquisa utilizando o hash do id
func BuscarPesquisaHash(Id_Hash string) model.Pesquisa {
	for _, pesquisa := range model.Pesquisas {
		if pesquisa.Id_Hash == Id_Hash {
			return pesquisa
		}
	}
	return model.Pesquisa{} // retorna uma pesquisa vazia
}

// ResponderPesquisaHash: responde a pesquisa utilizando o hash do id
func ResponderPesquisaHash(Id_Hash string, nota uint64) {
	for i, pesquisa := range model.Pesquisas {
		if pesquisa.Id_Hash == Id_Hash {
			model.Pesquisas[i].Nota = uint(nota)
		}
	}
}
