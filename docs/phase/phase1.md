# Phase 1: MVP開発 & デプロイメントフロー

本フェーズでは、「Backend実装 → Frontend実装 → コンテナ化(Docker) → GCPデプロイフロー確立」までの一連の流れを完遂し、OptiのMVPがクラウド上で稼働する状態を目指します。

## Step 1: Backend Implementation (Go & gRPC)

まずはバックエンドのマイクロサービス基盤と、コア機能（Catalog, User, etc.）を実装します。

### 1-1. プロジェクト基盤構築
最新のディレクトリ構成 (`proto` at root) に基づき初期化します。
- [ ] **Protobuf定義 (`proto/`)**
    - `proto/buf.yaml` の設定。
    - 各サービス (`product.proto`, `user.proto`) のメッセージ定義。
- [ ] **Go Workspace (`backend/go.work`)**
    - `backend/` 下にGo開発環境を集約。
    - `backend/gen/` へのコード生成設定 (`buf.gen.yaml`)。

### 1-2. マイクロサービス実装 (`backend/services/`)
各サービスを Clean Architecture で実装します。
- [ ] **Catalog Service (`services/catalog`)**
    - Firestore接続、製品データCRUDの実装。
- [ ] **User Service (`services/user`)**
    - ユーザー認証、コンテキスト管理。
- [ ] **Diagnosis / Adoption Service**
    - LLMロジック等はモックでも可とし、まずは通信疎通を優先。
- [ ] **動作確認**
    - `grpcurl` を用いてローカルでgRPCメソッドが叩けることを確認。

---

## Step 2: Frontend Implementation (Next.js)

バックエンドのAPIを利用するフロントエンドを実装します。

### 2-1. Next.js 基盤構築 (`frontend/`)
- [ ] **Next.js App Router 初期化**
- [ ] **Connect-Web Client 生成**
    - `proto/` から `frontend/src/gen` へTSコードを生成するフローを確立。

### 2-2. BFF & UI実装
- [ ] **BFF Layer (Server Components)**
    - 複数のマイクロサービスからデータを集約する Server Functions を実装。
- [ ] **UI Components**
    - 診断ウィザード画面、結果表示画面の実装。

---

## Step 3: Dockerization & Local Environment

それぞれのサービスをコンテナ化し、ローカルで本番に近い構成で動かします。

### 3-1. Dockerfile 作成 (`infra/docker/`)
- [ ] **Go Services**: マルチステージビルドを用いた軽量イメージの作成。
- [ ] **Next.js**: Standaloneモードを用いた実行イメージの作成。

### 3-2. ローカル実行環境 (`infra/compose/` or Skaffold)
- [ ] **Docker Compose 作成**
    - 全サービス + Emulator (Firestore) を一発で立ち上げる `docker-compose.yaml` を整備。
- [ ] **相互通信確認**
    - FrontendコンテナからBackendコンテナへgRPC通信が通ることを確認。

---

## Step 4: GCP Deployment Flow

ローカルで動いたコンテナを、Google Cloud Platform (GKE) へデプロイするフローを整えます。

### 4-1. インフラ準備 (Terraform 推奨)
- [ ] **GKE Cluster (Autopilot)** の構築。
- [ ] **Firestore (Native mode)** の有効化。
- [ ] **Artifact Registry** の作成。

### 4-2. デプロイメント確立
- [ ] **Build & Push**
    - コンテナイメージをArtifact RegistryへPush。
- [ ] **Manifests (`infra/k8s/`)**
    - Kubernetes Deployment/Service 定義ファイルの作成。
- [ ] **Apply & Verify**
    - `kubectl apply` によるデプロイ。
    - パブリックURL（またはLoadBalancer IP）からのアクセス確認。

---

**ゴール:**
これら全てが完了した時点で、Phase 1（= MVPリリース基盤の完成）とします。
