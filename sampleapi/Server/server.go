package server

import (
	"github.com/labstack/echo"
)

func Start() {
	e := echo.New()
	e.GET("/", landing)
	e.POST("/brabo", brabo)
	e.Logger.Fatal(e.Start(":8000"))
}
