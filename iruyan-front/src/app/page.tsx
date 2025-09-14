'use client';

import { useEffect } from "react";
import { useRouter } from "next/navigation";
import useUserStore from "@/stores/user-store";

function App() {
  const router = useRouter();
  const { currentUser } = useUserStore();

  useEffect(() => {
    // Zustandストアから現在のユーザー状態を確認
    if (currentUser) {
      // ユーザー情報がある場合は '/lobby' にリダイレクト
      router.replace("/lobby");
    } else {
      // ユーザー情報がない場合は '/login' にリダイレクト
      router.replace("/login");
    }
  }, [currentUser, router]);

  // リダイレクト中は何も表示しない（または必要に応じてローディングを表示）
  return null;
}

export default App;
