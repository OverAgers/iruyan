"use client";

import AuthorizationButton from "@/components/ui/button/main-button";
import AuthInputText from "@/components/ui/input/authorization-input-text";
import { useState } from "react";
import UseRegister from "../api/post-register";

export default function RegisterForm() {
  const [userId, setUserId] = useState("");
  const [password, setPassword] = useState("");
  const [userName, setUserName] = useState("");
  const [passwordConfirm, setPasswordConfirm] = useState("");
  const useRegister = UseRegister();

  const handleSubmit = () => {
    console.log(userId);
    console.log(password);
    console.log(userName);
    console.log(passwordConfirm);
    useRegister.setRequest({
      userID: userId,
      password: password,
      userName: userName,
    });
    console.log(useRegister.data);
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
          title="ユーザー名（表示名）"
          type="normal"
          setData={setUserName}
          placeholder="ユーザー名"
        />
        <AuthInputText
          title="パスワード"
          type="password"
          setData={setPassword}
          placeholder="パスワード"
        />
        <AuthInputText
          title="パスワード確認"
          type="password"
          setData={setPasswordConfirm}
          placeholder="パスワード確認"
        />
        <div>
          <AuthorizationButton title="新規登録" type="submit" />
        </div>
      </form>
    </div>
  );
}
