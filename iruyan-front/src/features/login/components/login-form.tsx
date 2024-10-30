"use client";

import MainButton from "@/components/ui/button/main-button";
import AuthInputText from "@/components/ui/input/authorization-input-text";

import UseLogin from "../api/post-login";
import { useState } from "react";

export default function LoginForm() {
  const [userName, setUserName] = useState("");
  const [password, setPassword] = useState("");
  const useLogin = UseLogin();

  const handleSubmit = () => {
    useLogin.login({ username: userName, password: password });
  };

  return (
    <div>
      <form className="flex-col">
        <AuthInputText
          title="ユーザーID"
          type="normal"
          setData={setUserName}
          placeholder="ユーザーID"
        />
        <AuthInputText
          title="パスワード"
          type="password"
          setData={setPassword}
          placeholder="パスワード"
        />
        <div>
          <AuthorizationButton title="入店する" type="submit" fullWidth={true} />
        </div>
      </form>
    </div>
  );
}
