'use client';

import React, {useEffect, useState} from "react";
import {apiFetch} from "../lib/apiClient";

export default function Home() {
  const [users, setUsers] = useState([]);
  useEffect(() => {
    apiFetch("/users").then(setUsers).catch(console.error);
  }, []);
  return (
    <main style={{padding:20}}>
      <h1>MyApp Frontend</h1>
      <p>Example: fetch users from Go backend</p>
      <ul>
        {users.map(u => <li key={u.id}>{u.name} (org:{u.org_id})</li>)}
      </ul>
    </main>
  );
}
