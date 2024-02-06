package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"plus-pesquisa/api/model"
	"plus-pesquisa/api/repository"
)

/*
BuscarFormPesquisa: função responsável por retornar um form para que o usuário responda a pesquisa
para essa prova de conceito irei apenas retornar apenas os dados da pesquisa para demonstração dos resultados esperados
*/
func BuscarFormPesquisa(c echo.Context) (err error) {
	idPesquisa, err := strconv.Atoi(c.Param("id")) //recuperação do parametro da url "id" e conversão dele para int
	if err != nil {                                //tratamento de erro do go.
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	pesquisa := repository.BuscarPesquisa(uint64(idPesquisa))
	return c.JSON(http.StatusOK, pesquisa)
}

// ResponderPesquisa: responde a pesquisa com uma nota
func ResponderPesquisa(c echo.Context) (err error) {
	idPesquisa, err := strconv.Atoi(c.Param("id")) //recuperação do parametro da url "id" e conversão dele para int
	if err != nil {                                //tratamento de erro do go.
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var pesquisaBodyRequisicao model.Pesquisa
	c.Bind(&pesquisaBodyRequisicao)

	repository.ResponderPesquisa(uint64(idPesquisa), uint64(pesquisaBodyRequisicao.Nota))

	return c.JSON(http.StatusOK, "pesquisa respondida com sucesso!")
}

func GerarLinkComIdHash(c echo.Context) (err error) {
	idPesquisa, err := strconv.Atoi(c.Param("id")) //recuperação do parametro da url "id" e conversão dele para int
	if err != nil {                                //tratamento de erro do go.
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	pesquisa := repository.BuscarPesquisa(uint64(idPesquisa))

	links := map[string]string{
		"Buscar Pesquisa":      fmt.Sprintf("http://localhost:8000/pesquisa/buscar/hash/%s", pesquisa.Id_Hash),
		"Buscar Form Pesquisa": fmt.Sprintf("http://localhost:8000/pesquisa/responder/hash/%s", pesquisa.Id_Hash),
	}

	return c.JSON(http.StatusOK, links)
}

/************************************Implementação das funções acima. porém com utilizando o hash do id ao invez do id *********************************************************/

func BuscarFormPesquisa_Hash(c echo.Context) (err error) {
	id_hash := c.Param("id_hash") //recuperação do parametro da url "id_hash"

	pesquisa := repository.BuscarPesquisaHash(id_hash)
	return c.JSON(http.StatusOK, pesquisa)
}

// ResponderPesquisa: responde a pesquisa com uma nota
func ResponderPesquisa_Hash(c echo.Context) (err error) {
	idPesquisa := c.Param("id_hash") //recuperação do parametro da url "id_hash"

	var pesquisaBodyRequisicao model.Pesquisa
	c.Bind(&pesquisaBodyRequisicao)

	repository.ResponderPesquisaHash(idPesquisa, uint64(pesquisaBodyRequisicao.Nota))

	return c.JSON(http.StatusOK, "pesquisa respondida com sucesso!")
}
