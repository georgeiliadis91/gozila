package controller

import (
	view "github.com/georgeiliadis91/gozila/views/layout"
	"github.com/labstack/echo/v4"
)

type IndexController struct {}

func (u IndexController) IndexControllerView(c echo.Context) error {
	return render(c, view.LayoutIndex())
}
