package repository

import (
	"context"
	"github.com/vshigimoto/LinkedIn-clone/internal/user/entity"
)

type Repository struct {
	UserRepo
}

type UserRepo interface {
	Create(ctx context.Context, user entity.User) (int, error)
	Get(ctx context.Context, userID int) (entity.User, error)
	Delete(ctx context.Context, userID int) error
	Update(ctx context.Context, user entity.User) error
}
