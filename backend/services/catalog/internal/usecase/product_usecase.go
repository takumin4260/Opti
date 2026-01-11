package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/kinoshitatakumi/opti/services/catalog/internal/domain/model"
	"github.com/kinoshitatakumi/opti/services/catalog/internal/domain/repository"
)

// ProductUsecase: 製品に関する「機能（ユースケース）」の実装です。
// API層（Handler）からリクエストを受け取り、ドメインモデルやリポジトリを使って
// 業務ロジック（ID生成、保存、検索など）を組み立てて実行します。
type ProductUsecase struct {
	repo repository.ProductRepository // データをどこに保存するかを知っている人（依存性注入）
}

// NewProductUsecase: ユースケースの作成
// リポジトリの実装を受け取ることで、保存先がメモリでもDBでも気にせず動くようになっています。
func NewProductUsecase(repo repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{repo: repo}
}

// CreateProduct: 製品作成のユースケース
// 1. IDがなければ生成する（業務ルール）
// 2. リポジトリを使って保存する
func (u *ProductUsecase) CreateProduct(ctx context.Context, input *model.Product) (*model.Product, error) {
	// 業務ロジック: IDの自動生成
	// クライアントがIDを指定してこなかった場合、ここでシステムがIDを割り振ります。
	if input.ID == "" {
		id, err := model.NewProductID(uuid.NewString())
		if err != nil {
			// ID生成に失敗（基本ありえないが、念のためエラーハンドリング）
			return nil, err
		}
		input.ID = id
	}

	// データの保存
	if err := u.repo.Save(ctx, input); err != nil {
		return nil, err
	}
	return input, nil
}

// ListProducts: 全製品取得のユースケース
// 特に複雑なロジックはなく、リポジトリからリストを取得してそのまま返します。
func (u *ProductUsecase) ListProducts(ctx context.Context) ([]*model.Product, error) {
	return u.repo.List(ctx)
}

// GetProduct: 指定したIDの製品を取得するユースケース
// stringで受け取ったIDを、安全な ProductID 型に変換してからリポジトリに渡します。
func (u *ProductUsecase) GetProduct(ctx context.Context, id string) (*model.Product, error) {
	// バリデーション: IDが空文字でないかチェック (NewProductIDの中で行われる)
	pid, err := model.NewProductID(id)
	if err != nil {
		return nil, err
	}
	return u.repo.GetByID(ctx, pid)
}
