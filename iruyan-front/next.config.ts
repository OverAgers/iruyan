import type { NextConfig } from "next";

// ---- Bundle Analyzer（ANALYZE=true のときだけ有効化）----
let withBundleAnalyzer: (cfg: NextConfig) => NextConfig = (cfg) => cfg;
if (process.env.ANALYZE === "true") {
  try {
    // NOTE: 三項演算で分岐しても良いが、try/catch で分かりやすく
    // TS/ESM でも Next が transpile するので require 可（だめなら下のコメント参照）
    // eslint-disable-next-line @typescript-eslint/no-var-requires, @typescript-eslint/no-require-imports
    withBundleAnalyzer = require("@next/bundle-analyzer")({ enabled: true });
    // もし ESM-only 環境で require が NG なら:
    // const mod = await import("@next/bundle-analyzer");
    // withBundleAnalyzer = mod.default({ enabled: true });
  } catch {
    console.warn(
      "ANALYZE=true ですが @next/bundle-analyzer が見つかりません。`pnpm add -D @next/bundle-analyzer` を実行してください。"
    );
  }
}

const nextConfig: NextConfig = {
  // ビルド中の ESLint を無視（必要に応じて）
  eslint: { ignoreDuringBuilds: true },

  // Docker の軽量ランタイムにするなら（任意）
  // output: "standalone",

  webpack: (config, { dev, isServer }) => {
    // 本番クライアントのみ分割戦略を上書き
    if (!dev && !isServer) {
      config.optimization.splitChunks = {
        chunks: "all",
        maxInitialRequests: 25,
        minSize: 20000,
        cacheGroups: {
          vendor: {
            test: /[\\/]node_modules[\\/]/,
            name: "vendors",
            priority: 10,
            reuseExistingChunk: true,
          },
          mui: {
            test: /[\\/]node_modules[\\/]@mui[\\/](?:material|system|base|icons|lab)[\\/]/,
            name: "mui",
            priority: 20,
            reuseExistingChunk: true,
          },
          // ※ ここが元コードだとパスにマッチしにくいので修正
          muiX: {
            test: /[\\/]node_modules[\\/]@mui[\\/]x-[^/\\]+[\\/]/,
            name: "mui-x",
            priority: 25,
            reuseExistingChunk: true,
          },
          emotion: {
            test: /[\\/]node_modules[\\/]@emotion[\\/]/,
            name: "emotion",
            priority: 15,
            reuseExistingChunk: true,
          },
          common: {
            name: "common",
            minChunks: 2,
            priority: 5,
            reuseExistingChunk: true,
          },
        },
      };
    }
    return config;
  },
};

export default withBundleAnalyzer(nextConfig);