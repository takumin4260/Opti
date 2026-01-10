package db

import (
	"context"
	"sync"

	"github.com/kinoshitatakumi/opti/services/catalog/internal/domain/model"
	"github.com/kinoshitatakumi/opti/services/catalog/internal/domain/repository"
)

// MemoryProductRepository: ドメイン層の Repository インターフェースを実装したクラスです。
// 現時点では本物のDB（PostgreSQLなど）の代わりに、メモリ上のマップ（変数）にデータを保存します。
// これにより、DBがなくても開発やテストを進めることができます。
type MemoryProductRepository struct {
	mu       sync.RWMutex                       // 排他制御用のロック（同時に書き込みが来た時に壊れないようにする）
	products map[model.ProductID]*model.Product // 実際のデータ保存場所（IDをキーにしたマップ）
}

// NewMemoryProductRepository: リポジトリの作成（初期化）を行います。
// 戻り値の型が interface (repository.ProductRepository) になっているのがポイントです。
// これにより、呼び出す側は「中身がメモリなのかDBなのか」を知らずに使えます。
func NewMemoryProductRepository() repository.ProductRepository {
	return &MemoryProductRepository{
		products: make(map[model.ProductID]*model.Product),
	}
}

// Save: 商品を保存（作成・更新）します。
func (r *MemoryProductRepository) Save(ctx context.Context, p *model.Product) error {
	r.mu.Lock()         // 書き込みロックを取得（他の人は読めない・書けない）
	defer r.mu.Unlock() // 関数が終わったら必ずアンロック
	r.products[p.ID] = p
	return nil
}

// List: 全ての商品リストを取得します。
func (r *MemoryProductRepository) List(ctx context.Context) ([]*model.Product, error) {
	r.mu.RLock()         // 読み取りロックを取得（他の人も読めるが、書き込めない）
	defer r.mu.RUnlock() // 関数が終わったら必ずアンロック

	// マップからスライス（配列）に変換して返します
	list := make([]*model.Product, 0, len(r.products))
	for _, p := range r.products {
		list = append(list, p)
	}
	return list, nil
}

// GetByID: IDを指定して商品を取得します。
func (r *MemoryProductRepository) GetByID(ctx context.Context, id model.ProductID) (*model.Product, error) {
	r.mu.RLock()         // 読み取りロック
	defer r.mu.RUnlock() // 必ずアンロック

	// マップに存在するかチェック
	if p, ok := r.products[id]; ok {
		return p, nil // 見つかったらポインタを返す
	}
	return nil, nil // 見つからなかったら nil を返す (エラーではない)
}
