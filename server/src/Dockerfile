# ビルドステージ
FROM golang:latest AS build

WORKDIR /go/src/github.com/Uwasaru/Sayka

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

# バイナリをビルド
RUN CGO_ENABLED=0 go build -o main main.go

# 最終ステージ
FROM alpine:latest

# ビルドステージからバイナリのみをコピー
COPY --from=build /go/src/github.com/Uwasaru/Sayka/main /app/main

# .envファイルをコピー
# COPY .env /app/.env

# コンテナ内の作業ディレクトリを設定
WORKDIR /app

# debug時はコメントアウト
CMD ["/app/main"]
