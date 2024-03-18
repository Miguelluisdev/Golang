package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func landing(c echo.Context) error {
	return c.String(http.StatusOK, "APi rodando")
}

func main() {
	e := echo.New()
	e.GET("/", landing)
	e.Logger.Fatal(e.Start(":8000"))
}
