package controller

import (
	view "github.com/georgeiliadis91/gozila/views"
	"github.com/labstack/echo/v4"
)

type StaticPageController struct {}


func (s StaticPageController) StaticPageControllerView(c echo.Context) error {
	return render(c, view.StaticPageTemplate())
}
