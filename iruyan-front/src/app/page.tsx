'use client';

import { useEffect } from "react";

function App() {
  useEffect(() => {
    // ユーザー情報を取得（例としてローカルストレージを使用）
    const userInfo = localStorage.getItem("user-store");

    if (userInfo) {
      // ユーザー情報がある場合は '/lobby' にリダイレクト
      window.location.href = "/lobby";
    } else {
      // ユーザー情報がない場合は '/login' にリダイレクト
      window.location.href = "/login";
    }
  }, []);

  return null; // 何も表示しない
}

export default App;
