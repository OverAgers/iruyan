"use client";

import AuthorizationButton from "@/components/ui/button/authorization-button";
import AuthInputText from "@/components/ui/input/authorization-input-text";

import UseLogin from "../api/post-login";
import { useState } from "react";

export default function LoginForm() {
  const [userId, setUserId] = useState("");
  const [password, setPassword] = useState("");
  const useLogin = UseLogin();

  const handleSubmit = () => {
    useLogin.setRequest({ userID: userId, password: password }); // リクエスト成功後にリダイレクト
    console.log(useLogin.data);
  };

  return (
    <div>
      <form className="flex-col" onSubmit={handleSubmit}>
        <AuthInputText
          title="ユーザーID"
          type="normal"
          setData={setUserId}
          placeholder="ユーザーID"
        />
        <AuthInputText
          title="パスワード"
          type="password"
          setData={setPassword}
          placeholder="パスワード"
        />
        <div>
          <AuthorizationButton title="入店する" type="submit" />
        </div>
      </form>
    </div>
  );
}
