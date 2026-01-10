# API設計書 (gRPC/Protobuf Definition Draft)

`docs/usecases.md` および `docs/domain_modeling.md` に基づくAPI定義案です。
バックエンド(Go)とフロントエンド(TypeScript)のインターフェース契約となります。

## 1. サービス一覧

| Service Name | 責務 | 関連ユースケース |
| :--- | :--- | :--- |
| `AuthService` | 認証・認可 (Signup, Login) | UC-00, UC-01 |
| `UserService` | ユーザー属性・住環境管理 | UC-01, Settings |
| `SimulationService` | 診断・シミュレーション実行 | UC-01, UC-02 |
| `ProjectService` | 採用プロジェクト・進捗管理 | UC-03, UC-04 |
| `ProductService` | 製品カタログ管理 (Admin) | UC-Admin-01 |

---

## 2. API詳細定義

### 2.1 AuthService / UserService
認証とユーザーコンテキスト（住環境）の管理。

```protobuf
service AuthService {
  // アカウント作成 (UC-01)
  rpc Signup(SignupRequest) returns (AuthResponse);
  // ログイン (UC-00)
  rpc Login(LoginRequest) returns (AuthResponse);
}

service UserService {
  // 住環境の更新 (UC-01, Settings)
  // 初回診断時もこれを使ってContextを作成・更新する
  rpc UpdateUserContext(UpdateUserContextRequest) returns (UserContext);
  
  // 住環境の取得
  rpc GetUserContext(GetUserContextRequest) returns (UserContext);
}

message UserContext {
  string id = 1;
  string user_id = 2;
  ResidenceInfo residence = 3;
}
```

### 2.2 SimulationService
診断と提案の実行。

```protobuf
service SimulationService {
  // シミュレーション実行 (UC-01, UC-02)
  // 入力された条件に基づき、提案(OptimizationPlan)を生成して返す
  rpc RunSimulation(RunSimulationRequest) returns (OptimizationPlan);

  // 過去のPlan詳細取得 (UC-02)
  rpc GetOptimizationPlan(GetOptimizationPlanRequest) returns (OptimizationPlan);

  // 自分のシミュレーション履歴取得 (Dashboard)
  rpc ListSimulationHistory(ListSimulationHistoryRequest) returns (ListSimulationHistoryResponse);
}

message RunSimulationRequest {
  // 住環境はUserContextから自動取得するか、一時的な上書きがあれば指定
  // ここでは「今回の条件」を指定
  BudgetConstraint budget = 1;
  repeated ChoreInput chores = 2;
}

message OptimizationPlan {
  string id = 1;
  string simulation_scenario_id = 2;
  string concept = 3;
  repeated ProposalGroup proposal_groups = 4;
  RoiProjection roi_projection = 5;
}
```

### 2.3 ProjectService
採用計画（AdoptionPlan）と実行進捗。

```protobuf
service ProjectService {
  // プロジェクト作成 (UC-03)
  // 提案プランから採用アイテムを選択して確定
  rpc CreateAdoptionProject(CreateAdoptionProjectRequest) returns (AdoptionPlan);

  // プロジェクト取得 (Dashboard, Detail)
  rpc GetAdoptionProject(GetAdoptionProjectRequest) returns (AdoptionPlan);
  rpc ListAdoptionProjects(ListAdoptionProjectsRequest) returns (ListAdoptionProjectsResponse);

  // アイテムステータス更新 (UC-04)
  // "bought", "installed" などの状態を変更
  rpc UpdateItemStatus(UpdateItemStatusRequest) returns (AdoptionPlan);
}

message CreateAdoptionProjectRequest {
  string source_plan_id = 1; // 元になったOptimizationPlan
  repeated SelectedItemSelection selections = 2; // どのアイテムを採用するか
}

message AdoptionPlan {
  string id = 1;
  string user_id = 2;
  ProjectStatus status = 3;
  repeated AdoptedItem items = 4;
}
```

### 2.4 ProductService (Admin)
製品カタログ管理。

```protobuf
service ProductService {
  // 製品一覧 (Admin)
  rpc ListProducts(ListProductsRequest) returns (ListProductsResponse);
  
  // 製品登録 (Admin)
  rpc CreateProduct(CreateProductRequest) returns (Product);
  
  // 製品更新 (Admin)
  rpc UpdateProduct(UpdateProductRequest) returns (Product);
  
  // 製品削除 (Admin)
  rpc DeleteProduct(DeleteProductRequest) returns (DeleteProductResponse);
}

message Product {
  string id = 1;
  string name = 2;
  string manufacturer = 3;
  int32 price = 4;
  string purchase_link = 5;
  // ...その他フィールド
}
```
