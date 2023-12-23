package controller

import (
	view "github.com/georgeiliadis91/gozila/views"
	"github.com/labstack/echo/v4"
)

type HomePageController struct {}

func (u HomePageController) HomePageControllerView(c echo.Context) error {
	return render(c, view.HomePageTemplate())
}
