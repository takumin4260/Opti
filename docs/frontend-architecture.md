# フロントエンド アーキテクチャ設計書

## 1. 概要
フロントエンドは **Next.js (App Router)** と **TypeScript** を使用して構築されます。
バックエンドサービスとは **Connect-Web** を用いて直接通信し、ブラウザから型安全なgRPC呼び出しを実現します。

## 2. 技術スタック
- **フレームワーク**: Next.js (App Router)
- **言語**: TypeScript
- **APIクライアント**: [Connect-ES](https://connect.build/docs/web/getting-started) (gRPC-Web / Connectプロトコル)
- **スタイリング**: Tailwind CSS
- **状態管理**: React Query (TanStack Query) - Connect-Query経由、または標準のHooks。

## 3. ディレクトリ構成

```text
frontend/
├── app/                      # App Router (BFF Layer for Data Fetching)
│   ├── (public)/             # LP, ログインなど
│   └── (dashboard)/          # Dashboard, 診断結果など (Private)
│
├── components/
│   ├── ui/                   # 汎用UIコンポーネント (Button, Card, Dialogなど)
│   ├── features/             # 機能単位のコンポーネント (Container/Presentational)
│   │   ├── diagnosis/
│   │   │   ├── DiagnosisWizard.tsx
│   │   │   └── ...
│   │   ├── catalog/
│   │   └── dashboard/
│   └── layout/               # ヘッダー, サイドバー, ラッパー
│
├── lib/
│   ├── rpc/                  # gRPC Client Factory
│   │   ├── server.ts         # [Server Side] Server Components/Actions用クライアント生成
│   │   └── client.ts         # [Client Side] Browser用クライアント生成
│   ├── actions/              # [Server Actions] BFF Mutation & Orchestration
│   │   ├── diagnosis.ts      # 例: ストリーミング診断実行, プラン保存
│   │   └── catalog.ts
│   └── utils/                # 共通ロジック
│
├── gen/                      # 生成されたProtobuf / Connectコード (Bufで管理)
│   └── opti/
│       └── ...
│
└── hooks/                    # Client Side Hooks
```

## 4. アーキテクチャパターン

### 4.1 バックエンドとの通信 & BFFパターン
**Next.js (Server Components) をBFF (Backend For Frontend) レイヤーとして位置付けます。**

*   **データ集約 (Aggregation)**:
    *   複数のマイクロサービス（例: Diagnosis結果 + Catalogの製品詳細）をまたぐデータ結合は、**Server Components** 内で行います。
    *   クライアント（ブラウザ）は、結合済みの描画データをPropsとして受け取るか、Server Actions経由で取得します。
*   **直接通信**:
    *   単純なデータ取得や、インタラクティブ性が高くServerを経由するオーバーヘッドを避けたい場合は、Client ComponentsからConnect-Webで直接サービスを叩くことも許容します。

### 4.2 クライアントコード生成
コードは `proto` 定義から `buf generate` コマンドを使用して生成されます。
生成されたコードは `frontend/gen` に配置され、Git管理に含めます（またはCIで生成）。

### 4.3 Feature-Sliced Design (軽量版)
コードベースの拡張性を保つため、すべてをAtomic Designにするのではなく、機能（Feature）ごとにロジックを集約します。
- `components/features/diagnosis/`: 診断フロー専用のコンポーネント群。
- `components/features/auth/`: ログイン・サインアップ関連。

## 5. 開発ガイドライン

### Proto統合フロー
1. ルートの `proto` ディレクトリにある `.proto` ファイルを定義・更新する。
2. `buf generate` を実行し、`frontend/gen` 内のTypeScriptクライアントコードを更新する。
3. コンポーネントでService Clientをインポートして使用する。

```typescript
// 使用例
import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { UserCommandService } from "@/gen/opti/user/v1/user_connect";

const transport = createConnectTransport({
  baseUrl: "https://api.opti.local",
});
const client = createClient(UserCommandService, transport);

await client.updateProfile({ ... });
```

### スタイリング
- すべてのスタイリングに **Tailwind CSS** を使用します。
- ランタイムオーバーヘッドを避けるため、CSS-in-JSライブラリの使用は避けます。
- 条件付きクラス名には `clsx` または `tailwind-merge` を使用します。
