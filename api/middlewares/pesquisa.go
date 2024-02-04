package middlewares

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"plus-pesquisa/api/repository"
)

/*
	O middleware é uma camada que recebe a resquisição antes ou depois do controller. depende da forma que irá ser posicionado.
*/

/*
O middleware VerificarSePesquisaRespondida, irá verificar se a pesquisa já foi respondida. em caso afirmativo irá retornar a
mensagem "está pesquisa já foi respondida", caso contrário a requisição irá para o controller onde será processada.
*/
func VerificarSePesquisaRespondida(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		idPesquisa, err := strconv.Atoi(c.Param("id")) //recuperação do parametro da url "id" e conversão dele para int
		if err != nil {                                //tratamento de erro do go.
			fmt.Println(err)
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		if repository.VerificarSePesquisaRespondida(uint64(idPesquisa)) {
			return c.JSON(http.StatusForbidden, "está pesquisa já foi respondida")
		}

		return next(c)
	}
}
