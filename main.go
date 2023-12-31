package main

import (
	"backend-kakojapan/controller"
	"backend-kakojapan/db"
	"backend-kakojapan/repository"
	"backend-kakojapan/router"
	"backend-kakojapan/usecase"
)

func main() {
	db := db.NewDB()
	imageRepository := repository.NewImageRepository(db)
	imageUsecase := usecase.NewImageUsecase(imageRepository)
	imageController := controller.NewImageController(imageUsecase)
	e := router.NewRouter(imageController)
	e.Logger.Fatal(e.Start(":8080"))
}
