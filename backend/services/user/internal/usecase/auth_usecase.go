package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/kinoshitatakumi/opti/services/user/internal/domain/model"
	"github.com/kinoshitatakumi/opti/services/user/internal/domain/repository"
)

type AuthUsecase struct {
	repo repository.UserRepository
}

func NewAuthUsecase(repo repository.UserRepository) *AuthUsecase {
	return &AuthUsecase{repo: repo}
}

func (u *AuthUsecase) SignUp(ctx context.Context, input *model.User) (*model.User, error) {
	if input.ID == "" {
		id := uuid.NewString()
		input.ID = id
	}

	if err := u.repo.Save(ctx, input); err != nil {
		return nil, err
	}
	return input, nil
}

func (u *AuthUsecase) Login(ctx context.Context, input *model.User) (*model.User, error) {
	return u.repo.GetByEmail(ctx, input.Email)
}