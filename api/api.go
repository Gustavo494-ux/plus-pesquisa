package api

import (
	"fmt"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"

	"plus-pesquisa/api/rotas"
)

const (
	PortaApi = 8000
)

// Configurar: configura a API para utilização
func ConfigurarApi() {
	e := rotas.GerarEcho()

	StartarApi(e, time.Second*30)
}

// StartarApi: realiza o start da api
func StartarApi(e *echo.Echo, timeout time.Duration) {
	e.Server.WriteTimeout = timeout
	if err := e.Start(fmt.Sprintf(":%s", strconv.Itoa(PortaApi))); err != nil {
		fmt.Println(err)
	}
}
