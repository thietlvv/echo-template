package usecase

import (
	"context"
)

type UsersUsecase interface {
	GetAllUsers(ctx context.Context) error
}
