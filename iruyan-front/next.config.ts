import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  // 他の設定オプションがある場合はここに追加
  eslint: {
    ignoreDuringBuilds: true, // ビルド中のESLintエラーを無視する設定
  },
};

export default nextConfig;
