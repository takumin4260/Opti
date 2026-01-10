package db

import (
	"context"
	"fmt"

	"github.com/kinoshitatakumi/opti/services/catalog/internal/domain/model"
	"github.com/kinoshitatakumi/opti/services/catalog/internal/domain/repository"
)

// FirestoreProductRepository: Firestoreを使用したリポジトリ実装
type FirestoreProductRepository struct {
	client *FirestoreClient
}

// NewFirestoreProductRepository: コンストラクタ
// クライアント(接続)を注入してもらうことで、責務を分離しています。
func NewFirestoreProductRepository(client *FirestoreClient) repository.ProductRepository {
	return &FirestoreProductRepository{client: client}
}

const collectionName = "products"

func (r *FirestoreProductRepository) Save(ctx context.Context, p *model.Product) error {
	// Firestore独自の保存処理
	// map[string]interface{} への変換が必要な場合もありますが、
	// firestoreタグをつければ直接保存も可能です。
	_, err := r.client.Client.Collection(collectionName).Doc(p.ID.String()).Set(ctx, p)
	if err != nil {
		return fmt.Errorf("failed to save product to firestore: %w", err)
	}
	return nil
}

// List: 仮実装
func (r *FirestoreProductRepository) List(ctx context.Context) ([]*model.Product, error) {
	// 実装イメージ:
	// docs, err := r.client.Client.Collection(collectionName).Documents(ctx).GetAll()
	return nil, fmt.Errorf("not implemented yet")
}

// GetByID: 仮実装
func (r *FirestoreProductRepository) GetByID(ctx context.Context, id model.ProductID) (*model.Product, error) {
	doc, err := r.client.Client.Collection(collectionName).Doc(id.String()).Get(ctx)
	if err != nil {
		return nil, err
	}
	var p model.Product
	if err := doc.DataTo(&p); err != nil {
		return nil, err
	}
	return &p, nil
}
