#!/bin/bash

# このスクリプトはマイグレーションを実行するためのものです
# 使用例:
#   ./migrate.sh         # マイグレーションを実行
#   ./migrate.sh -r      # マイグレーションをロールバック

set -e

# 引数の処理
ROLLBACK=""
if [ "$1" = "-r" ] || [ "$1" = "--rollback" ]; then
  ROLLBACK="--rollback"
  echo "マイグレーションロールバックを実行します..."
else
  echo "マイグレーションを実行します..."
fi

# マイグレーションの実行
docker compose exec api /app/migrate $ROLLBACK

echo "完了しました"
