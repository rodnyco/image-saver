package photo

import (
	"context"
	"rodny/image-saver/models"
)

type Repository interface {
	CreatePhoto(ctx context.Context, photo *models.Photo) error
}
