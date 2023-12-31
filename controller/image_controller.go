package controller

import (
	"backend-kakojapan/model"
	"backend-kakojapan/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IImageController interface {
	CreateImage(c echo.Context) error
}

type imageController struct {
	iu usecase.IImageUsecase
}

func NewImageController(iu usecase.IImageUsecase) IImageController {
	return &imageController{iu}
}

func (ic *imageController) CreateImage(c echo.Context) error {
	image := model.Image{}
	if err := c.Bind(&image); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	imageRes, err := ic.iu.CreateImage(image)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, imageRes)
}
