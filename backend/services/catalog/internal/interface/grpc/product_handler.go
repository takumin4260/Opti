package grpc

import (
	"context"

	"connectrpc.com/connect"
	catalogv1 "github.com/kinoshitatakumi/opti/gen/go/catalog/v1"
	"github.com/kinoshitatakumi/opti/pkg/domain/value"
	"github.com/kinoshitatakumi/opti/services/catalog/internal/domain/model"
	"github.com/kinoshitatakumi/opti/services/catalog/internal/usecase"
)

// ProductHandler: gRPCリクエストを受け付ける「窓口」です。
// Clean Architectureにおける「Interface層」にあたります。
// 外部からの通信(gRPC)と、内部のロジック(Usecase)の通訳を行います。
type ProductHandler struct {
	usecase *usecase.ProductUsecase // 実際の処理を行う人（依存性注入）
}

// NewProductHandler: ハンドラの作成
func NewProductHandler(u *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{usecase: u}
}

// ListProducts: 製品一覧取得API
func (h *ProductHandler) ListProducts(ctx context.Context, req *connect.Request[catalogv1.ListProductsRequest]) (*connect.Response[catalogv1.ListProductsResponse], error) {
	// 1. ユースケースを呼び出してデータを取得（内部の型 model.Product が返ってくる）
	products, err := h.usecase.ListProducts(ctx)
	if err != nil {
		return nil, err
	}

	// 2. 内部の型(model) -> 通信用(protobuf) に変換
	var pbProducts []*catalogv1.Product
	for _, p := range products {
		pbProducts = append(pbProducts, &catalogv1.Product{
			Id:                     p.ID.String(),
			Name:                   p.Name,
			Description:            p.Description,
			Price:                  p.Price.Amount(),
			Manufacturer:           p.Manufacturer,
			PurchaseLink:           p.PurchaseLink,
			ImageUrl:               p.ImageURL,
			WeakPoints:             p.WeakPoints,
			StrongPoints:           p.StrongPoints,
			InstallationDifficulty: string(p.InstallationDifficulty),
			Category:               string(p.Category),
		})
	}

	return connect.NewResponse(&catalogv1.ListProductsResponse{
		Products: pbProducts,
	}), nil
}

// CreateProduct: 製品作成API
func (h *ProductHandler) CreateProduct(ctx context.Context, req *connect.Request[catalogv1.CreateProductRequest]) (*connect.Response[catalogv1.Product], error) {
	// 1. バリデーション: 通信用の型(int32) -> 内部の値オブジェクト(Price) に変換
	// ここで「マイナス価格」などの不正な値を弾きます。
	price, err := value.NewPrice(req.Msg.Price)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	// 2. 通信用(protobuf) -> 内部の型(model) に変換
	input := &model.Product{
		Name:                 req.Msg.Name,
		Description:            req.Msg.Description,
		Price:                  price,
		Manufacturer:           req.Msg.Manufacturer,
		PurchaseLink:           req.Msg.PurchaseLink,
		ImageURL:               req.Msg.ImageUrl,
		WeakPoints:             req.Msg.WeakPoints,
		StrongPoints:           req.Msg.StrongPoints,
		InstallationDifficulty: model.InstallationDifficulty(req.Msg.InstallationDifficulty),
		Category:               model.ProductCategory(req.Msg.Category),
	}

	p, err := h.usecase.CreateProduct(ctx, input)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&catalogv1.Product{
		Id:                     p.ID.String(),
		Name:                   p.Name,
		Description:            p.Description,
		Price:                  p.Price.Amount(),
		Manufacturer:           p.Manufacturer,
		PurchaseLink:           p.PurchaseLink,
		ImageUrl:               p.ImageURL,
		WeakPoints:             p.WeakPoints,
		StrongPoints:           p.StrongPoints,
		InstallationDifficulty: string(p.InstallationDifficulty),
		Category:               string(p.Category),
	}), nil
}

// GetProduct: 製品詳細取得API
func (h *ProductHandler) GetProduct(ctx context.Context, req *connect.Request[catalogv1.GetProductRequest]) (*connect.Response[catalogv1.Product], error) {
	// 1. ユースケースを呼び出す
	p, err := h.usecase.GetProduct(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	}
	// 2. 存在しない場合は 404 NotFound エラーを返す
	if p == nil {
		return nil, connect.NewError(connect.CodeNotFound, nil)
	}

	// 3. 内部の型(model) -> 通信用(protobuf) に変換してレスポンス

	return connect.NewResponse(&catalogv1.Product{
		Id:                     p.ID.String(),
		Name:                   p.Name,
		Description:            p.Description,
		Price:                  p.Price.Amount(),
		Manufacturer:           p.Manufacturer,
		PurchaseLink:           p.PurchaseLink,
		ImageUrl:               p.ImageURL,
		WeakPoints:             p.WeakPoints,
		StrongPoints:           p.StrongPoints,
		InstallationDifficulty: string(p.InstallationDifficulty),
		Category:               string(p.Category),
	}), nil
}

func (h *ProductHandler) UpdateProduct(ctx context.Context, req *connect.Request[catalogv1.UpdateProductRequest]) (*connect.Response[catalogv1.Product], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, nil) // 未実装
}

func (h *ProductHandler) DeleteProduct(ctx context.Context, req *connect.Request[catalogv1.DeleteProductRequest]) (*connect.Response[catalogv1.DeleteProductResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, nil) // 未実装
}
