package router

import (
	"backend-kakojapan/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(ic controller.IImageController) *echo.Echo {
	e := echo.New()
	e.POST("/image", ic.CreateImage)
	return e
}
