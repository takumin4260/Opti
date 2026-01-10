# 全体ディレクトリ構成 (Monorepo Directory Structure)

`backend-architecture.md`, `frontend-architecture.md` を統合した、リポジトリ全体のディレクトリ構成定義です。

## ルート構成
```text
Opti/ (Project Root)
├── docs/                     # ドキュメント
│
├── infra/                    # [Infrastructure] Docker, Terraform, CI/CD
│
├── proto/                    # [Schema] API定義 (Shared Source of Truth)
│   ├── catalog/v1/
│   ├── user/v1/
│   └── buf.yaml              # Buf Module Config
│
├── frontend/                 # [Frontend] Next.js Application
│   ├── src/
│   │   ├── gen/              # 生成されたTSクライアント (from ../proto)
│   └── ...
│
└── backend/                  # [Backend] Go Microservices
    ├── go.work               # Go Workspace
    ├── gen/                  # 生成されたGoインターフェース (from ../proto)
    │   └── go/
    └── services/             # マイクロサービス実装
        ├── user/
        ├── catalog/
        └── ...
```

---

## 各ディレクトリ詳細

### 1. `proto/` (Shared Schema)
**API定義の正本**。FrontendとBackendの間の「契約」であり、特定の言語に依存しないため、ルート直下に配置するのがベストプラクティスです。
*   ここにある `.proto` ファイルから、Frontend用(TS)とBackend用(Go)のコードをそれぞれのディレクトリに生成します。

### 2. `infra/` (Infrastructure)
Docker Compose, Terraform, Cloud Build設定など、アプリケーションコード以外のインフラ構成要素をここに集約します。
*   `tools` よりも具体的で標準的な命名です。

### 3. `backend/` (Go Monorepo)
*   **`backend/gen/`**: `proto/` から生成されたGoコードの出力先。
*   **`backend/go.work`**: Backend開発のルート。

### 4. `frontend/` (Next.js)
*   **`frontend/src/gen/`**: `proto/` から生成されたTSコードの出力先。

---

## 修正点
ユーザー様のご指摘に基づき、**Schema (proto)** と **Infrastructure** を第一級市民としてルートに配置しました。
これにより、「FrontendとBackendは対等であり、共通のSchema（proto）に依存する」という構造が明確になります。

