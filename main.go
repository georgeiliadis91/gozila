package main

import (
	"github.com/georgeiliadis91/gozila/controller"
	"github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()
    i := controller.IndexController{}
	e.Static("/assets", "assets")
    e.GET("/", i.IndexControllerView)
    e.Logger.Fatal(e.Start(":3000"))
}
