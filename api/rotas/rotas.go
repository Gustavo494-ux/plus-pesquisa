package rotas

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// GerarEcho: retorna uma instância de Echo configurada inclusive as rotas
func GerarEcho() *echo.Echo {
	e := echo.New()

	configurarRotas(e)

	return e
}

// configurarRotas: realiza a configuração das rotas
func configurarRotas(e *echo.Echo) {
	e.GET("/teste", func(c echo.Context) error {
		return c.String(http.StatusOK, time.Now().Format(time.RFC3339Nano))
	})

	RotasPesquisa(e)
}
