package repoository

import (
	"context"
	"schedule-remind-go/entity"
)

type UserRepository interface {
	Store(ctx context.Context, usr *entity.User) (*entity.User, error)
	UserGetByID(ctx context.Context, lineId string) (*entity.User, error)
}
