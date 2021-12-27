package photo

import "context"

type UseCase interface {
	CreatePhoto(ctx context.Context, photoCodeBase64 string) (string, error)
}
