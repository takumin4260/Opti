package model

import "fmt"

// ProductID: 製品ID専用の型 (Value Object) です。
// 単なる string ではなく、専用の型を作ることで、ユーザーIDや他のIDと混同するミスを防ぎます。
// 例: func GetProduct(id ProductID) は、誤って UserID を渡すとコンパイルエラーになります。
type ProductID string

// NewProductID: 安全に ProductID を生成するコンストラクタです。
// 「空文字はIDとして認めない」というルール（バリデーション）をここで保証します。
func NewProductID(value string) (ProductID, error) {
	if value == "" {
		return "", fmt.Errorf("ProductID cannot be empty")
	}
	return ProductID(value), nil
}

// String: 元の文字列値を取り出すメソッドです。
// DB保存やAPIレスポンス生成時に、string型に戻すために使います。
func (id ProductID) String() string {
	return string(id)
}
