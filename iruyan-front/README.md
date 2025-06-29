# Iruyan Frontend

Iruyanのフロントエンドアプリケーションです。Next.js 15を使用して構築されています。

## 技術スタック

- **Framework**: Next.js 15 (App Router)
- **Language**: TypeScript
- **Package Manager**: pnpm
- **Styling**: Tailwind CSS + MUI
- **State Management**: Zustand
- **Form Handling**: React Hook Form + Zod
- **HTTP Client**: Axios
- **Data Fetching**: SWR

## 開発環境のセットアップ

### 前提条件

- Node.js 18.0.0以上
- pnpm 8.0.0以上

### インストール

```bash
# 依存関係のインストール
pnpm install

# 環境変数の設定
cp env.example .env.local
# .env.localファイルを編集して必要な環境変数を設定
```

### 開発サーバーの起動

```bash
# 開発サーバー起動
pnpm dev

# 型チェック
pnpm type-check

# リント実行
pnpm lint

# リント修正
pnpm lint:fix

# ビルド
pnpm build

# 本番サーバー起動
pnpm start

# キャッシュクリア
pnpm clean
```

## プロジェクト構造

```
src/
├── app/                 # Next.js App Router
├── components/          # 再利用可能なコンポーネント
│   ├── ui/             # 基本UIコンポーネント
│   ├── layouts/        # レイアウトコンポーネント
│   └── providers/      # プロバイダーコンポーネント
├── features/           # 機能別モジュール
│   ├── auth/          # 認証機能
│   ├── lobby/         # ロビー機能
│   ├── rooms/         # ルーム機能
│   └── ...
├── hooks/             # カスタムフック
├── lib/               # ライブラリ設定
├── stores/            # Zustandストア
├── types/             # TypeScript型定義
├── utils/             # ユーティリティ関数
└── styles/            # グローバルスタイル
```

## パスエイリアス

以下のパスエイリアスが設定されています：

- `@/*` - srcディレクトリ
- `@/components/*` - コンポーネント
- `@/features/*` - 機能モジュール
- `@/hooks/*` - カスタムフック
- `@/lib/*` - ライブラリ設定
- `@/stores/*` - ストア
- `@/types/*` - 型定義
- `@/utils/*` - ユーティリティ
- `@/styles/*` - スタイル

## 環境変数

必要な環境変数は `env.example` ファイルを参照してください。

## Docker

```bash
# イメージビルド
docker build -t iruyan-front .

# コンテナ起動
docker run -p 3000:3000 iruyan-front
```

## 開発ガイドライン

1. **型安全性**: TypeScriptを活用し、anyの使用を避ける
2. **コンポーネント設計**: 再利用可能で保守しやすいコンポーネントを作成
3. **エラーハンドリング**: 適切なエラーハンドリングを実装
4. **パフォーマンス**: 不要な再レンダリングを避ける
5. **アクセシビリティ**: アクセシビリティを考慮したUI設計

## ライセンス

このプロジェクトはMITライセンスの下で公開されています。

## Getting Started

First, run the development server:

```bash
npm run dev
# or
yarn dev
# or
pnpm dev
# or
bun dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

You can start editing the page by modifying `app/page.tsx`. The page auto-updates as you edit the file.

This project uses [`next/font`](https://nextjs.org/docs/app/building-your-application/optimizing/fonts) to automatically optimize and load [Geist](https://vercel.com/font), a new font family for Vercel.

## Learn More

To learn more about Next.js, take a look at the following resources:

- [Next.js Documentation](https://nextjs.org/docs) - learn about Next.js features and API.
- [Learn Next.js](https://nextjs.org/learn) - an interactive Next.js tutorial.

You can check out [the Next.js GitHub repository](https://github.com/vercel/next.js) - your feedback and contributions are welcome!

## Deploy on Vercel

The easiest way to deploy your Next.js app is to use the [Vercel Platform](https://vercel.com/new?utm_medium=default-template&filter=next.js&utm_source=create-next-app&utm_campaign=create-next-app-readme) from the creators of Next.js.

Check out our [Next.js deployment documentation](https://nextjs.org/docs/app/building-your-application/deploying) for more details.
