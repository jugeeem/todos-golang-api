# Todos Golang API

TodosアプリケーションのバックエンドAPIです。ユーザー認証とTodoタスクの管理機能を提供します。

## 機能

- ユーザー登録・ログイン（JWT認証）
- Todoタスクの作成・取得・更新・削除
- ユーザーごとのTodoタスク管理

## 技術スタック

- 言語: Go 1.23
- フレームワーク: Gin
- データベース: PostgreSQL
- ORM: GORM
- マイグレーション: golang-migrate
- 認証: JWT

## セットアップ

### 必要条件

- Docker および Docker Compose
- Go 1.23以上

### 環境変数

`.env`ファイルを作成し、以下の環境変数を設定してください:

```
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=todo_db
DB_SSLMODE=disable
JWT_SECRET_KEY=your_secret_key
BCRYPT_COST_FACTOR=12
PORT=8080
```

### 実行方法

1. リポジトリをクローン:

```bash
git clone https://github.com/jugeeem/golang-todo.git
cd golang-todo
```

2. アプリケーションを起動:

```bash
docker-compose up -d
```

3. マイグレーションを実行:

```bash
./migrate.sh
```

## API エンドポイント

### 認証

- `POST /api/v1/register` - ユーザー登録
- `POST /api/v1/token` - ログイン (JWTトークン取得)

### ユーザー

- `GET /api/v1/users` - 全ユーザー取得
- `GET /api/v1/users/:id` - 特定ユーザー取得
- `PUT /api/v1/users/:id` - ユーザー情報更新
- `DELETE /api/v1/users/:id` - ユーザー削除

### Todo

- `GET /api/v1/todos` - 全Todoタスク取得
- `POST /api/v1/todos` - 新規Todoタスク作成
- `GET /api/v1/todos/:id` - 特定のTodoタスク取得
- `PUT /api/v1/todos/:id` - Todoタスク更新
- `DELETE /api/v1/todos/:id` - Todoタスク削除
- `GET /api/v1/todos/my` - ログインユーザーのTodoタスク取得

## プロジェクト構成

```
/
├── app/
│   ├── domain/         - ドメインモデルとリポジトリインターフェース
│   ├── infrastructure/ - データベース・マイグレーション関連
│   ├── interface/      - HTTPハンドラとルーティング
│   ├── usecase/        - ビジネスロジック
│   ├── utility/        - ユーティリティ関数
│   └── main.go         - アプリケーションエントリーポイント
├── cmd/
│   └── migrate/        - マイグレーション実行用コマンド
├── migrations/         - データベースマイグレーションファイル
├── docker-compose.yml  - Docker構成ファイル
├── Dockerfile          - Dockerビルド設定
├── go.mod              - Goモジュール定義
└── migrate.sh          - マイグレーション実行スクリプト
```

## ライセンス

MIT License - 詳細は[LICENSE](LICENSE)ファイルを参照してください。