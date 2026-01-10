# 実装フェーズ計画 (Implementation Phases)

小さな機能単位で垂直（Backend〜Frontend疎通）に立ち上げていくための段階的実装プランです。

| Phase | テーマ | 主な実装内容 | ゴール |
| :--- | :--- | :--- | :--- |
| **Phase 1** | **プロジェクト基盤 & 製品カタログ** | プロジェクト構成、Protobuf基盤、Catalog Service (Read) | Admin画面で製品データが表示・管理できる |
| **Phase 2** | **ユーザーコンテキスト** | User Service, 認証連携, 住環境データ保存 | 診断フォームでユーザー情報が保存・取得できる |
| **Phase 3** | **シミュレーション (Core)** | Simulation Service, 提案ロジック実装 | 診断を実行し、AI提案結果が表示される |
| **Phase 4** | **採用・実行 (Execution)** | Project Service, 進捗管理機能 | ダッシュボードで導入計画のステータス管理ができる |

---

詳細なタスクは各Phaseのドキュメントを参照してください。
*   [Phase 1 詳細: 基盤構築 & カタログ機能](./phase1.md)
*   Phase 2, 3, 4 (TBD)
