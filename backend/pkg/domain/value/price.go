package value

import "fmt"

// Price: 金額を表す値オブジェクト
// マイナス値を許容しないなどの不変条件を保証します。
type Price struct {
	amount   int32
	currency string
}

// NewPrice: Priceのコンストラクタ
// バリデーションを行い、不正な値の場合はエラーを返します。
func NewPrice(amount int32) (Price, error) {
	if amount < 0 {
		return Price{}, fmt.Errorf("price cannot be negative: %d", amount)
	}
	// MVPではJPY固定とします
	return Price{
		amount:   amount,
		currency: "JPY",
	}, nil
}

// Amount: 金額を取得するためのゲッター
func (p Price) Amount() int32 {
	return p.amount
}

// Currency: 通貨を取得するためのゲッター
func (p Price) Currency() string {
	return p.currency
}
