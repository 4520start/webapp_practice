export async function apiFetch(path, options = {}) {
  // const baseUrl = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";

  const isServer = typeof window === "undefined";
  
  let baseUrl;
  if (isServer) {
    // サーバサイドではコンテナ名を使用
    baseUrl = process.env.NEXT_PUBLIC_API_URL || "http://backend:8080";
  } else {
    // クライアント（ブラウザ）ではホストのポート番号を使用
    baseUrl = `${window.location.protocol}//${window.location.hostname}:8080`;
  }
  

  const res = await fetch(`${baseUrl}${path}`, {
    ...options,
    headers: {
      "Content-Type": "application/json",
      ...(options.headers || {}),
    },
  });
  if (!res.ok) throw new Error(`API Error: ${res.status}`);
  return res.json();
}
