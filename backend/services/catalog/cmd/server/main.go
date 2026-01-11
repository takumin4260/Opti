package main

import (
	"log"
	"net/http"

	"github.com/kinoshitatakumi/opti/gen/go/catalog/v1/catalogv1connect"
	"github.com/kinoshitatakumi/opti/services/catalog/internal/infrastructure/db"
	"github.com/kinoshitatakumi/opti/services/catalog/internal/interface/grpc"
	"github.com/kinoshitatakumi/opti/services/catalog/internal/usecase"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	// 1. Dependency Injection (依存性の注入)
	// ここでアプリケーションの全ての部品を生成し、組み立てます。

	// (a) Repository: データの保存場所を作成（今回はメモリ）
	// 将来ここを `db.NewPostgresProductRepository()` に変えるだけでDB切り替えが完了します。
	repo := db.NewMemoryProductRepository()

	// (b) Usecase: ビジネスロジックを作成
	// 作成したリポジトリを渡すことで、Useaseは保存場所を知らずに使えます。
	u := usecase.NewProductUsecase(repo)

	// (c) Handler: 外部との窓口を作成
	// 作成したUseaseを渡すことで、リクエストをロジックに流せるようにします。
	handler := grpc.NewProductHandler(u)

	// 2. サーバーのルーティング設定
	mux := http.NewServeMux()
	// Connectが生成したコードを使って、「このパスに来たら、このハンドラを呼ぶ」という紐付けを行います。
	path, connectHandler := catalogv1connect.NewProductServiceHandler(handler)
	mux.Handle(path, connectHandler)

	// 3. サーバー起動
	log.Println("Starting catalog service on :8080")
	// h2c: HTTP/2 Cleartext (暗号化なしHTTP/2)
	// gRPCは通常HTTP/2が必要ですが、開発中はTLS(SSL)設定が面倒なので、
	// 平文でHTTP/2を喋れる `h2c` を使って起動します。
	err := http.ListenAndServe(":8080", h2c.NewHandler(mux, &http2.Server{}))
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
