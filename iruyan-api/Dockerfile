# ベースイメージを指定
FROM golang:1.23-alpine

# コンテナ内の作業ディレクトリを設定
WORKDIR /app

# 必要なツールをインストール
RUN apk add --no-cache git

# air と swagのインストール
RUN go install github.com/cosmtrek/air@v1.40.4 \
  && go install github.com/swaggo/swag/cmd/swag@v1.8.12

# ローカルのソースコードをコンテナにコピー
COPY . .

# 必要なパッケージをインストール
RUN go mod download

# アプリケーションをビルド
RUN go build -o main .

# 公開予定のコンテナのポートを明示
EXPOSE 8080

# PATHにairとswagのパスを追加
ENV PATH=$PATH:/go/bin

# アプリケーションを実行
CMD ["air"]
