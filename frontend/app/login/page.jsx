'use client';

import React, { useState } from "react";
import { useRouter } from "next/navigation";
import { apiFetch } from "../../lib/apiClient";

export default function LoginPage() {
  const [isRegister, setIsRegister] = useState(false);
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const router = useRouter();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError("");
    
    try {
      const endpoint = isRegister ? "/register" : "/login";
      const result = await apiFetch(endpoint, {
        method: "POST",
        body: JSON.stringify({ username, password }),
      });
      
      // ログイン成功：セッション保存
      localStorage.setItem("user", JSON.stringify(result));
      router.push("/");
    } catch (err) {
      setError(isRegister ? "登録に失敗しました" : "ログインに失敗しました");
    }
  };

  return (
    <main style={{ padding: 40, maxWidth: 400, margin: "0 auto" }}>
      <h1>{isRegister ? "ユーザー登録" : "ログイン"}</h1>
      <form onSubmit={handleSubmit}>
        <div style={{ marginBottom: 16 }}>
          <label>ユーザー名</label>
          <input
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
            style={{ width: "100%", padding: 8 }}
          />
        </div>
        <div style={{ marginBottom: 16 }}>
          <label>パスワード</label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
            style={{ width: "100%", padding: 8 }}
          />
        </div>
        {error && <p style={{ color: "red" }}>{error}</p>}
        <button type="submit" style={{ padding: "8px 16px" }}>
          {isRegister ? "登録" : "ログイン"}
        </button>
      </form>
      <p style={{ marginTop: 16 }}>
        <button onClick={() => setIsRegister(!isRegister)} style={{ background: "none", border: "none", color: "blue", cursor: "pointer" }}>
          {isRegister ? "ログインに切り替え" : "新規登録に切り替え"}
        </button>
      </p>
    </main>
  );
}