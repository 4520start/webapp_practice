'use client';

import React, { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { apiFetch } from "../lib/apiClient";

export default function Home() {
  const [users, setUsers] = useState([]);
  const [currentUser, setCurrentUser] = useState(null);
  const router = useRouter();

  useEffect(() => {
    // 認証チェック
    const userStr = localStorage.getItem("user");
    if (!userStr) {
      router.push("/login");
      return;
    }
    setCurrentUser(JSON.parse(userStr));
    
    // ユーザー一覧取得
    apiFetch("/users").then(setUsers).catch(console.error);
  }, [router]);

  const handleLogout = () => {
    localStorage.removeItem("user");
    router.push("/login");
  };

  if (!currentUser) return <p>Loading...</p>;

  return (
    <main style={{ padding: 20 }}>
      <div style={{ display: "flex", justifyContent: "space-between", alignItems: "center" }}>
        <h1>MyApp Frontend</h1>
        <div>
          <span style={{ marginRight: 16 }}>ようこそ、{currentUser.username} さん</span>
          <button onClick={handleLogout}>ログアウト</button>
        </div>
      </div>
      <p>Example: fetch users from Go backend</p>
      <ul>
        {users.map(u => <li key={u.id}>{u.name} (org:{u.org_id})</li>)}
      </ul>
    </main>
  );
}



// 'use client';

// import React, {useEffect, useState} from "react";
// import {apiFetch} from "../lib/apiClient";

// export default function Home() {
//   const [users, setUsers] = useState([]);
//   useEffect(() => {
//     apiFetch("/users").then(setUsers).catch(console.error);
//   }, []);
//   return (
//     <main style={{padding:20}}>
//       <h1>MyApp Frontend</h1>
//       <p>Example: fetch users from Go backend</p>
//       <ul>
//         {users.map(u => <li key={u.id}>{u.name} (org:{u.org_id})</li>)}
//       </ul>
//     </main>
//   );
// }