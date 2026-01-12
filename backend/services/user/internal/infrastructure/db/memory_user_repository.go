package db

import (
	"context"
	"errors"
	"sync"

	"github.com/kinoshitatakumi/opti/pkg/domain/value"
	"github.com/kinoshitatakumi/opti/services/user/internal/domain/model"
	"github.com/kinoshitatakumi/opti/services/user/internal/domain/repository"
)

// MemoryUserRepository は UserRepository のインメモリ実装です。
// サーバーを再起動するとデータは消えます。
type MemoryUserRepository struct {
	mu       sync.RWMutex
	users    map[string]*model.User
	contexts map[string]*model.UserContext
}

// NewMemoryUserRepository は新しい MemoryUserRepository を作成します。
func NewMemoryUserRepository() repository.UserRepository {
	return &MemoryUserRepository{
		users:    make(map[string]*model.User),
		contexts: make(map[string]*model.UserContext),
	}
}

// Save はユーザーを保存します。
func (r *MemoryUserRepository) Save(ctx context.Context, user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// 既存ユーザーの上書きも許可するシンプルな実装
	r.users[user.ID] = user
	return nil
}

// GetByEmail はEmailでユーザーを検索します。
func (r *MemoryUserRepository) GetByEmail(ctx context.Context, email value.Email) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// MapにはEmailのインデックスがないので、全件走査します（DBならIndexが効く場面）
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, errors.New("user not found")
}

// GetUserContext はユーザーIDでコンテキストを取得します。
func (r *MemoryUserRepository) GetUserContext(ctx context.Context, userID string) (*model.UserContext, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	ctxData, ok := r.contexts[userID]
	if !ok {
		return nil, errors.New("user context not found")
	}
	return ctxData, nil
}

// SaveUserContext はコンテキストを保存/更新します。
func (r *MemoryUserRepository) SaveUserContext(ctx context.Context, userCtx *model.UserContext) error {
	if userCtx.UserID == "" {
		return errors.New("invalid user context: UserID is required")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	// UserIDをキーとして保存
	r.contexts[userCtx.UserID] = userCtx
	return nil
}
