package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func landing(c echo.Context) error {
	return c.String(http.StatusOK, "API rodando")
}

type Message struct {
	Msg     string   `json:"msg"`
	Code    string   `json:"code"`
	Hobbies []string `json:"Hobbies" `
}

func brabo(c echo.Context) error {
	code := c.FormValue("code")
	res := Message{}
	res.Code = code
	res.Msg = "API do golang teste"
	res.Hobbies = []string{"basquete", "volei", "futebol"}
	return c.JSON(http.StatusOK, res)
}
