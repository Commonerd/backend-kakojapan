package usecase

import (
	"backend-kakojapan/model"
	"backend-kakojapan/repository"
)

// 인터페이스 정의
type IImageUsecase interface {
	CreateImage(image model.Image) (model.ImageResponse, error)
}

// 유즈케이스 구조체 정의
type imageUsecase struct {
	ir repository.IImageRepository
}

func NewImageUsecase(ir repository.IImageRepository) IImageUsecase {
	return &imageUsecase{ir}
}

// 이미지 생성 함수 정의
func (iu *imageUsecase) CreateImage(image model.Image) (model.ImageResponse, error) {
	//새로운 이미지 생성
	newImage := model.Image{Title: image.Title, Description: image.Description, Lat: image.Lat, Lng: image.Lng, Date: image.Date}

	// 이미지 생성 및 에러 처리
	if err := iu.ir.CreateImage(&newImage); err != nil {
		return model.ImageResponse{}, err
	}

	// 이미지 응답 모델 생성
	resImage := model.ImageResponse{
		ID:    newImage.ID,
		Title: newImage.Title,
	}

	// 응답모델 반환
	return resImage, nil

}
