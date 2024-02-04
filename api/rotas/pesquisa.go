package rotas

import (
	"github.com/labstack/echo/v4"

	"plus-pesquisa/api/controller"
	"plus-pesquisa/api/middlewares"
)

// RotasPesquisa: Criação das rotas de pesquisa
func RotasPesquisa(e *echo.Echo) {
	pesquisa := e.Group("/pesquisa") //Grupo de rotas da pesquisa

	// Está rota não possui um middleware que impede seu processamento caso a pesquisa já esteja respondida
	pesquisa.GET("/buscar/:id", controller.BuscarFormPesquisa)

	// as rotas abaixo possuem um middleware que impede seu processamento caso a pesquisa já esteja respondida
	pesquisa.GET("/responder/:id", controller.BuscarFormPesquisa, middlewares.VerificarSePesquisaRespondida)
	pesquisa.PUT("/responder/:id", controller.ResponderPesquisa, middlewares.VerificarSePesquisaRespondida)
}
