package model

import "github.com/kinoshitatakumi/opti/pkg/domain/value"

// Product: ドメインモデルとしての製品定義です。
// Protoファイル(通信用)とは異なり、Goのプログラム内でビジネスロジックを扱うための純粋な構造体です。
// 特定のライブラリ（DBタグやJSONタグなど）に依存させないことで、技術的な変更に強くしています。
type Product struct {
	ID                     ProductID              // システム内で一意なID (Value Object)
	Name                   string                 // 製品名
	Description            string                 // 製品の詳細説明
	Price                  value.Price            // 価格 (Value Object)
	Manufacturer           string                 // 製造メーカー名
	PurchaseLink           string                 // 購入サイトへのURL
	ImageURL               string                 // 製品画像のURL
	WeakPoints             []string               // 導入時のデメリット・注意点 (AIによる分析用)
	StrongPoints           []string               // 導入時のメリット・アピールポイント
	InstallationDifficulty InstallationDifficulty // 設置難易度
	Category               ProductCategory        // 製品カテゴリ
}

// InstallationDifficulty: 設置難易度を表す型
type InstallationDifficulty string

const (
	DifficultyLow    InstallationDifficulty = "low"
	DifficultyMedium InstallationDifficulty = "medium"
	DifficultyHigh   InstallationDifficulty = "high"
)

// ProductCategory: 製品カテゴリを表す型
type ProductCategory string

const (
	CategoryRobotVacuum ProductCategory = "robot_vacuum"
	CategorySmartLock   ProductCategory = "smart_lock"
	CategoryDishWasher  ProductCategory = "dishwasher"
	CategoryLighting    ProductCategory = "lighting"
	CategorySensor      ProductCategory = "sensor"
	CategoryHub         ProductCategory = "hub"
	CategoryOther       ProductCategory = "other"
)
