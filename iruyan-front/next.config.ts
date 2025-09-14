import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  // 他の設定オプションがある場合はここに追加
  eslint: {
    ignoreDuringBuilds: true, // ビルド中のESLintエラーを無視する設定
  },
  // webpack設定
  webpack: (config, { dev, isServer }) => {
    // 本番環境でのバンドルサイズ最適化
    if (!dev && !isServer) {
      config.optimization.splitChunks = {
        chunks: 'all',
        maxInitialRequests: 25,
        minSize: 20000,
        cacheGroups: {
          vendor: {
            test: /[\\/]node_modules[\\/]/,
            name: 'vendors',
            priority: 10,
            reuseExistingChunk: true,
          },
          mui: {
            test: /[\\/]node_modules[\\/]@mui[\\/]/,
            name: 'mui',
            priority: 20,
            reuseExistingChunk: true,
          },
          muiX: {
            test: /[\\/]node_modules[\\/]@mui\/x-[\\/]/,
            name: 'mui-x',
            priority: 25,
            reuseExistingChunk: true,
          },
          emotion: {
            test: /[\\/]node_modules[\\/]@emotion[\\/]/,
            name: 'emotion',
            priority: 15,
            reuseExistingChunk: true,
          },
          common: {
            name: 'common',
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

// Bundle Analyzerの設定
const withBundleAnalyzer = require('@next/bundle-analyzer')({
  enabled: process.env.ANALYZE === 'true',
});

export default withBundleAnalyzer(nextConfig);
