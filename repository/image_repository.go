package repository

import (
	"backend-kakojapan/model"

	"gorm.io/gorm"
)

type IImageRepository interface {
	//GetImageById(image *model.Image, id string) error
	CreateImage(image *model.Image) error
}

type imageRepository struct {
	db *gorm.DB
}

func NewImageRepository(db *gorm.DB) IImageRepository {
	return &imageRepository{db}
}

func (ir *imageRepository) CreateImage(image *model.Image) error {
	if err := ir.db.Create(image).Error; err != nil {
		return err
	}
	return nil
}
