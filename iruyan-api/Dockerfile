# ベースイメージを指定
FROM golang:1.23-alpine

# コンテナ内の作業ディレクトリを設定
WORKDIR /app

# ローカルのソースコードをコンテナにコピー
COPY . .

# 必要なパッケージをインストール
RUN go mod download

# アプリケーションをビルド
RUN go build -o main .

# 公開予定のコンテナのポートを明示
EXPOSE 8080

# ホットリロードツール air をインストール
RUN go install github.com/air-verse/air@latest

# アプリケーションを実行
CMD ["air"]