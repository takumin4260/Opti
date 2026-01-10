# ユースケース & 画面遷移 (Draft)

MVP要件とドメインモデルに基づき、ユーザーの行動フローと必要な画面を定義します。

## 1. コア・ユースケース (User Journeys)

### UC-01: 初回診断とアカウント作成
ユーザーが初めてサービスを訪れ、現状を入力して最初の「最適解」を見るまで。
1. LPでサービスの価値（「家事の消滅」）を理解する。
2. アカウント登録なしで簡易診断を開始する（あるいは最初から登録）。
3. **住環境**（UserContext）と**現在の悩み**（Pains）を入力する。
4. **予算**（Simulation input）を設定する。
5. 提案結果（OptimizationPlan）が表示され、その内容を保存するためにアカウント登録を行う。

### UC-02: シミュレーションの再試行 (条件変更)
提案されたプランに対し、条件を変えて再計算する。
1. 提案結果画面で「予算をもっと抑えたい」「やっぱりこの悩みは優先しなくていい」と入力を修正する。
2. 新しい条件で再計算（Create new SimulationScenario）を実行する。
3. 新しいプランが表示される。履歴として残る。

### UC-03: 導入ロードマップの確認と実行
提案内容に納得し、実際に購入・導入を進める。
1. 提案された「今月買うもの」リスト（Roadmap）を確認する。
2. 各製品の詳細・購入リンクへ飛ぶ。
3. （Optional: 購入完了をステータスに反映する -> MVP範囲外かもしれないが、あると便利）

---

## 2. ページ構成案 (Sitemap)

### Public
*   `/`: ランディングページ (LP)
    *   FV: キャッチコピー
    *   Start Button: 診断開始

### Onboarding / Diagnosis (Wizard形式)
*   `/diagnosis/step1`: 住環境入力 (戸建て/マンション, 間取り, 段差など)
*   `/diagnosis/step2`: 生活スタイル・ペイン入力 (家事時間, 苦痛度)
*   `/diagnosis/step3`: 予算・優先度設定
*   `/diagnosis/analyzing`: ローディング演出 (「AIが最適解を計算中...」)

### Result & Simulation (Core)
*   `/plan/[planId]`: 提案結果メイン
    *   **Concept**: AIによる提案コンセプト (Concept)
    *   **ROI**: 時間削減効果、投資回収イメージ
    *   **Structure**: 掃除・洗濯などカテゴリごとの提案 (ProposalGroups)
    *   **Action**: 条件変更して再計算ボタン

*   `/plan/[planId]/roadmap`: 導入順序チャート
    *   STEP1, STEP2... の時系列表示

### Product Detail
*   `/products/[productId]`: 製品詳細 (Modal or Page)
    *   スペック、メリット、購入リンク

### User (My Page)
*   `/dashboard`: 過去のシミュレーション履歴一覧
*   `/settings/residence`: 住環境基本設定の修正
*   `/settings/account`: アカウント設定
