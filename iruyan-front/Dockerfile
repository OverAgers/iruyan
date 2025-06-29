# ベースイメージとしてNode.js 22.xを使用
FROM node:22-alpine

# pnpmをグローバルにインストール（package.jsonのpackageManagerと一致）
RUN npm install -g pnpm@10.12.4

# 作業ディレクトリを作成
WORKDIR /app

# package.jsonとpnpm-lock.yamlをコピー
COPY package.json pnpm-lock.yaml ./

# 依存パッケージのインストール
RUN pnpm install --frozen-lockfile

# アプリケーションのソースコードをすべてコピー
COPY . .

# Next.jsアプリのビルド（本番モード用）
RUN pnpm build

# Next.jsアプリの起動ポート
EXPOSE 3000

# コンテナ起動時にNext.jsのアプリを起動（本番モード）
CMD ["pnpm", "start"]
