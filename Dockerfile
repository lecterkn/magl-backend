FROM golang:1.24.1

WORKDIR /app

# 依存関係取得
COPY go.mod go.sum ./
RUN go mod download

# プロジェクトをコピー
COPY ./migrations ./migrations
COPY ./docs/ ./docs
COPY ./internal/ ./internal
COPY ./cmd/ ./cmd
COPY ./dbconfig.yml .

# ビルド
RUN CGO_ENABLED=0 GOOS=linux go build -o ./go_app ./cmd/main.go

# マイグレーション
RUN go install github.com/rubenv/sql-migrate/...@latest

# 実行
ENTRYPOINT [ "sh", "-c", "sql-migrate up -env=\"production\" && ./go_app" ]
