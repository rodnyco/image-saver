package usecase

import (
	"context"
	"encoding/base64"
	"rodny/image-saver/pkg/file"
)

type PhotoUseCase struct {
	//photoRepo photo.Repository
}

func NewPhotoUseCase() *PhotoUseCase {
	return &PhotoUseCase{}
}

func (photo PhotoUseCase) CreatePhoto(ctx context.Context, photoCodeBase64 string) (string, error) {
	decode, err := base64.StdEncoding.DecodeString(photoCodeBase64)
	if err != nil {
		return "", err
	}

	fileName := file.CreateFileName()
	path, err := file.SaveFile("./images", fileName + ".png", decode)
	if err != nil {
		return "", err
	}

	return path, nil
	//newPhoto := &models.Photo{
	//	Name: fileName,
	//	Path: path,
	//}

	//return photo.photoRepo.CreatePhoto(ctx, newPhoto)
}
