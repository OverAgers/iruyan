# ベースイメージとしてNode.js 20.12.2を使用
FROM node:20.12.2

# 作業ディレクトリを作成
WORKDIR /app

# package.jsonとyarn.lockをコピー
COPY package.json yarn.lock ./

# 依存パッケージのインストール
RUN yarn install

# アプリケーションのソースコードをすべてコピー
COPY . .

# Next.jsアプリのビルド（本番モード用）
RUN yarn build

# Next.jsアプリの起動ポート
EXPOSE 3000

# コンテナ起動時にNext.jsのアプリを起動（本番モード）
CMD ["yarn", "start"]
