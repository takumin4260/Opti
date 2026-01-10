package db

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
)

// FirestoreClient: Firestoreへの接続を管理するクライアント
type FirestoreClient struct {
	Client *firestore.Client
}

// NewFirestoreClient: Firestoreクライアントの初期化
func NewFirestoreClient(ctx context.Context, projectID string) (*FirestoreClient, error) {
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to create firestore client: %w", err)
	}
	return &FirestoreClient{Client: client}, nil
}

// Close: 接続を閉じる
func (c *FirestoreClient) Close() error {
	return c.Client.Close()
}
