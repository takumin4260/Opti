package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/kinoshitatakumi/opti/services/user/internal/domain/model"
	"github.com/kinoshitatakumi/opti/services/user/internal/domain/repository"
)

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) GetUserContext(ctx context.Context, userID string) (*model.UserContext, error) {
	return u.repo.GetUserContext(ctx, userID)
}

func (u *UserUsecase) SaveUserContext(ctx context.Context, userID string, context *model.UserContext) error {
	if context.ID == "" {
		id := uuid.NewString()
		context.ID = id
	}
	context.UserID = userID
	return u.repo.SaveUserContext(ctx, context)
}