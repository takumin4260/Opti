# バックエンド アーキテクチャ設計書

## 1. 概要
Optiのバックエンドシステムは、**Go言語** で実装された **マイクロサービスアーキテクチャ** を採用しています。
サービス間およびフロントエンドとの通信には **gRPC (Connect)** を使用します。
各サービスの内部設計は、保守性とテスト容易性を担保するため、**クリーンアーキテクチャ** および **ドメイン駆動設計 (DDD)** の原則に従います。

## 2. 技術スタック
- **言語**: Go
- **RPCフレームワーク**: [Connect](https://connect.build/) (gRPC互換。HTTP/1.1 & HTTP/2をサポート)
- **データベース**: NoSQL (Firestore)
- **DI (依存性の注入)**: [Google Wire](https://github.com/google/wire)

## 3. マイクロサービス構成 (Service Boundaries)

`docs/api_design.md` の定義に基づき、以下の4つのマイクロサービスに分割します。

| Service Directory | Service Name | 担当領域 (Context) | API Definition |
| :--- | :--- | :--- | :--- |
| `services/user` | **User Service** | 認証 (Auth), ユーザー属性, 住環境 (User Context) | `AuthService`, `UserService` |
| `services/catalog` | **Catalog Service** | 製品データベース管理 (Product Context) | `ProductService` |
| `services/simulation` | **Simulation Service** | 診断ロジック, 提案生成 (Recommendation Context) | `SimulationService` |
| `services/project` | **Project Service** | 採用プラン, 進捗管理 (Execution Context) | `ProjectService` |

---

## 4. ディレクトリ構成 (Monorepo Layout)

プロジェクトルート直下に `services/` ディレクトリを配置し、各サービスを格納します。

```text
Opti/
├── gen/                      # 自動生成コード (Protobuf -> Go/TS)
├── proto/                    # Protobuf定義ファイル (.proto)
│
├── services/                 # マイクロサービス群
│   ├── user/                 # [User Service]
│   │   ├── cmd/server/       # Entrypoint
│   │   ├── internal/         # Clean Architecture (Domain, Usecase...)
│   │   └── go.mod
│   │
│   ├── catalog/              # [Catalog Service]
│   │   └── ...
│   │
│   ├── simulation/           # [Simulation Service]
│   │   └── ...
│   │
│   └── project/              # [Project Service]
│       └── ...
│
└── go.work                   # Go Workspace (複数モジュール管理)
```

### 各サービスの内部構成 (Standard Go Layout + Clean Arch)
各サービス内部 (`services/xxx/`) は、共通して以下の構造を持ちます。

```text
service-name/
├── cmd/
│   └── server/
│       └── main.go           # エントリーポイント。依存関係を解決しサーバーを起動。
├── internal/
│   ├── domain/               # [Inner Layer]
│   │   ├── model/            # エンティティ定義 (DDD)
│   │   └── repository/       # Repositoryインターフェース
│   ├── usecase/              # [Application Layer]
│   ├── interface/            # [Adapter Layer] (gRPC Handlers, DB Gateway)
│   └── infrastructure/       # [Framework Layer] (DB Config, Logger)
└── go.mod
```

## 4. レイヤーの責務

### 4.1 Domain層 (`internal/domain`)
- **役割**: ビジネスロジックとルールの中核を表現します。**外部依存を一切持ちません。**
- **コンポーネント**:
    - **Models**: ふるまいを持つ構造体（例: `User`, `Recommendation`）。
    - **Repository Interfaces**: データの「取得方法（How）」を定義せず、「何ができるか（What）」のみを定義します（例: `UserRepository`）。

### 4.2 Usecase層 (`internal/usecase`)
- **役割**: ドメインオブジェクトを組み合わせて、特定のアプリケーション要件（ユースケース）を達成します。
- **コンポーネント**:
    - **Interactors**: Interface層からデータを受け取り、Domainモデルを操作して、Repository経由で保存します。
    - **Input/Output Ports**: (Goでは暗黙的になることも多いですが) 入出力の境界を定義します。

### 4.3 Interface層 (`internal/interface`)
- **役割**: 外部形式のデータを内部形式に変換、またはその逆を行います。
- **コンポーネント**:
    - **gRPC Handlers**: 自動生成されたProtobufインターフェースを実装します。ProtoメッセージをDomain/Usecaseモデルに変換します。
    - **Gateways (Repositories)**: Domain層で定義されたRepositoryインターフェースを、実際のドライバ（Firestore等）を使って実装します。

### 4.4 Infrastructure層 (`internal/infrastructure`)
- **役割**: 技術的な詳細設定やフレームワークへの依存を閉じ込めます。
- **コンポーネント**:
    - DB接続の初期化。
    - ロガーの設定。
    - 環境変数のパース。

## 5. 実装ガイドライン

### 依存性の注入 (Dependency Injection)
依存の方向は常に **内側（Domain）** に向かいます。
- `Infrastructure` は `Interface` を知っている。
- `Interface` は `Usecase` を知っている。
- `Usecase` は `Domain` を知っている。
- `Domain` は外側のレイヤーを一切知らない。

`cmd/server/main.go` での依存関係解決コード生成には **Wire** を使用します。

### エラーハンドリング
- Domain/Usecase層は、標準的なGoの `error` または独自のドメインエラー型を返すべきです。
- gRPC Handler層が、これらのエラーをRPCステータスコード（例: `connect.NewError(connect.CodeNotFound, err)`）に変換する責任を持ちます。

### データベース (NoSQL Mapping)
- Domainモデルは、極力DBタグ（`firestore:"..."` 等）を持たず、技術的詳細から独立させます。
- **データ参照**: `Diagnosis` ドメイン内では `Product` エンティティそのものではなく、`ProductId` (Value Object) を保持します。
    - アプリケーション層（Usecase）で詳細情報が必要になった場合のみ、Catalog Service経由で解決（Resolve）します。
