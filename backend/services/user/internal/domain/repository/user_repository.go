package repository

import (
	"context"

	"github.com/kinoshitatakumi/opti/pkg/domain/value"
	"github.com/kinoshitatakumi/opti/services/user/internal/domain/model"
)

type UserRepository interface {
	Save(ctx context.Context, user *model.User) error
	GetByEmail(ctx context.Context, email value.Email) (*model.User, error)
	GetUserContext(ctx context.Context, userID string) (*model.UserContext, error)
	SaveUserContext(ctx context.Context, context *model.UserContext) error
}
