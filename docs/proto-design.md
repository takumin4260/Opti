# プロトコルバッファ設計書 & レジストリ

## 1. 概要
Optiでは、すべてのサービスにおけるインターフェース記述言語 (IDL) として **Protocol Buffers (Protobuf)** を使用します。
これが、Frontend-Backend間、およびBackendマイクロサービス間におけるAPI契約（コントラクト）の唯一の正解情報となります。

## 2. 管理戦略
- **ツール**: [Buf](https://buf.build/)
- **リポジトリ**: モノレポ。すべての `.proto` ファイルはプロジェクトルートの `proto/` ディレクトリ配下に置きます。
- **Linting (静的解析)**: `buf lint` を使用して厳格に強制します。
- **破壊的変更の検知**: `buf breaking` を使用して監視します。

## 3. パッケージ構成
互換性を担保するため、バージョニングされたパッケージ構成を採用します。

```text
proto/
├── opti/
│   ├── user/
│   │   └── v1/
│   │       ├── user.proto
│   │       └── service.proto
│   ├── catalog/
│   │   └── v1/
│   │       ├── product.proto
│   │       └── service.proto
│   └── diagnosis/
│       └── v1/
│           ├── diagnosis.proto
│           └── service.proto
└── buf.yaml                        # ワークスペース設定
```

### 名前空間の規則
`package opti.<service>.v1;`

例: `package opti.user.v1;`

## 4. サービス定義 (ドラフト)

### 4.1 User Service (`opti.user.v1`)
ユーザープロファイル管理および住環境データに焦点を当てます。

**RPCs**:
- `GetMyProfile(Empty) returns (UserProfile)`
- `UpdateHousingProfile(HousingProfile) returns (UserProfile)`
- `UpdateLifeStyle(LifeStyle) returns (UserProfile)`

**Messages**:
- `UserProfile`: 全ユーザー情報の集約。
- `HousingProfile`: 住居形態、間取り、制約事項。
- `LifeStyle`: 起床・就寝時間、ペインポイント。

### 4.2 Catalog Service (`opti.catalog.v1`)
製品情報のマスター管理・検索を行います。

**RPCs**:
- `GetProduct(GetProductRequest) returns (Product)`
- `ListProductsByIds(ListProductsByIdsRequest) returns (ListProductsResponse)`: 提案詳細表示時に使用
- `SearchProducts(SearchProductsRequest) returns (ListProductsResponse)`: 名前やカテゴリでの検索

**Messages**:
- `Product`: 製品のマスターデータ（名前、カテゴリ、価格、設置難易度など）。

### 4.3 Diagnosis Service (`opti.diagnosis.v1`)
ユーザープロファイルに基づいたレコメンド生成を受け持ちます。
製品情報は内部にコピーせず、`product_id` のみを参照として保持します。

**RPCs**:
- `CreateSimulation(CreateSimulationRequest) returns (SimulationScenario)`: 診断セッション開始
- `GeneratePlan(GeneratePlanRequest) returns (OptimizationPlan)`: LLMによるプラン生成
- `GetPlan(GetPlanRequest) returns (OptimizationPlan)`: プラン詳細取得

**Messages**:
- `SimulationScenario`: 診断コンテキスト（UserContextのスナップショットを含む）。
- `OptimizationPlan`: 提案グループ (`ProposalGroup`) を持つ最終成果物。
- `ProposedItem`: `product_id` と `quantity`, `reason` のみを持つ軽量な参照ポインタ。

## 5. スタイルガイド
- **フィールド名**: `snake_case` (Protobuf標準)。
- **RPC名**: `PascalCase` (例: `GetProfile`).
- **Enum名**: `UPPER_SNAKE_CASE` で、プレフィックスにサービス名を付ける (例: `HOUSING_TYPE_APARTMENT`).
- **Null許容性**: 標準のラッパー型を使用するか、特定の「未定義」状態が必要な場合は `optional` キーワードを使用します（ただしgRPCはゼロ値を基本とします）。
