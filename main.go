package main

import (
	controller "github.com/georgeiliadis91/gozila/controllers"
	"github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()
    i := controller.HomePageController{}
    s := controller.StaticPageController{}
	e.Static("/assets", "assets")
    e.GET("/", i.HomePageControllerView)
    e.GET("/:static_page", s.StaticPageControllerView)
    e.Logger.Fatal(e.Start(":3000"))
}
