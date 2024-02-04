package rotas

import (
	"github.com/labstack/echo/v4"

	"plus-pesquisa/api/controller"
)

// RotasPesquisa: Criação das rotas de pesquisa
func RotasPesquisa(e *echo.Echo) {
	pesquisa := e.Group("/pesquisa") //Grupo de rotas da pesquisa

	pesquisa.GET("/responder/:id", controller.BuscarFormPesquisa)
}
