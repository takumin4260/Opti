# ユースケース & 画面遷移 (Draft)

MVP要件とドメインモデルに基づき、ユーザーの行動フローと必要な画面を定義します。

## 1. コア・ユースケース (User Journeys)

### UC-00: 再来訪と活動再開 (Returning User)
既存ユーザーがサイトに戻り、活動を再開する。
1. **ログイン**: アカウント情報でログインする。
2. **ダッシュボード確認**: 現在進行中のプロジェクト（`AdoptionPlan`）の進捗状況を確認する。
3. **活動選択**: 「プロジェクトのタスクを進める」か「新しい条件で別のシミュレーションを試す」かを選択する。

### UC-01: 住環境の登録と初期診断 (UserContext & First Simulation)
ユーザーがサービスを利用開始し、自分の環境を登録する。
1. **住環境入力**: 居住タイプ、間取り、物理制約などを入力し、`UserContext`を作成する。
2. **悩み・予算入力**: 現在の家事の悩みと予算を入力する。
3. **シミュレーション実行**: `SimulationScenario`を作成し、即座に`OptimizationPlan`（提案）を受け取る。

### UC-02: 条件を変えて再シミュレーション (SimulationScenario)
提案されたプランを調整する。
1. **条件変更**: 「予算を5万円下げたい」「掃除よりも洗濯を優先したい」等の調整を行う。
2. **再計算**: 新しい`SimulationScenario`を作成（住環境はスナップショットとして保持）。
3. **比較**: 過去のシミュレーション結果と比較検討する（Dashboardで履歴確認）。

### UC-03: プランの採用とプロジェクト化 (Adoption)
提案内容から、実際に導入するものを決定する（Selection Phase）。
1. **アイテム選別**: 提案された`OptimizationPlan`の中から、実際に導入する製品にチェックを入れる。
2. **プロジェクト作成**: 選んだアイテムで`AdoptionPlan`（導入プロジェクト）を作成・確定する。
   *   これ以降、このプランは「ただの提案」から「実行計画」に変わる。

### UC-04: 導入進捗の管理 (Execution)
作成したプロジェクトの進捗を追跡する。
1. **購入**: 各製品のリンクから購入し、ステータスを「購入済み(bought)」に変更する。
2. **設置**: 届いた製品を設置し、ステータスを「設置完了(installed)」にする。
3. **完了**: 全てのアイテムが設置完了になると、プロジェクト完了（ゴール達成）。

---

## 2. 管理者ユースケース (Admin Use Cases)

### UC-Admin-01: 製品カタログ管理 (Product Management)
ユーザーに提案するためのIoT製品データベースを管理する。
1. **製品登録**: 新しい製品のスペック、価格、タグ（features）、紹介リンク等を登録する。
2. **製品更新**: 価格変動や新モデル発売に伴い、情報を更新する。
3. **製品削除**: 販売終了した製品をアーカイブまたは削除する。

---

## 3. ページ構成案 (Sitemap)

### Public
*   `/`: ランディングページ (LP)
*   `/login`: ログインページ (Email/Password)
    *   FV: キャッチコピー
    *   Start Button: 診断開始

### Onboarding / Diagnosis (Wizard形式)
*   `/diagnosis/residence`: 住環境入力 (UserContext作成)
*   `/diagnosis/lifestyle`: 生活スタイル・ペイン入力 (Simulation Input)
*   `/diagnosis/budget`: 予算・優先度設定 (Simulation Input)
*   `/diagnosis/analyzing`: ローディング演出

### Simulation Result (Proposal Phase)
*   `/plans/[planId]`: 提案結果メイン (OptimizationPlan)
    *   **Structure**: カテゴリ別提案 (ProposalGroups)
    *   **Action**: 
        *   「条件を変えて再計算」 (-> New Simulation)
        *   「このプランで採用へ進む」 (-> Adoption Flow)

### Adoption / Project (Execution Phase)
*   `/projects/create?planId=xxx`: 採用アイテム選択画面
    *   提案アイテムの一覧からCheckboxで選択して確定
*   `/projects/[adoptionId]`: 進行中の導入プロジェクト詳細
    *   **Trello/Kanban風リスト**: 「未購入」「購入済み」「設置完了」のステータス管理
    *   **Action**: ステータス更新、メモ書き込み

### User (My Page)
*   `/dashboard`: 
    *   進行中のプロジェクト (`AdoptionPlan`)
    *   過去のシミュレーション履歴 (`SimulationScenario` list)
*   `/settings/residence`: 住環境基本設定の修正 (`UserContext` update)
*   `/settings/account`: アカウント設定

### Admin (Internal)
*   `/admin/dashboard`: 管理機能トップ
*   `/admin/products`: 製品一覧
*   `/admin/products/new`: 製品登録
*   `/admin/products/[productId]/edit`: 製品編集
