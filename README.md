# MyApp Sample (Next.js + Go + Postgres)

## 構成
- frontend: Next.js (simple) - ports: 3000
- backend: Go (Echo + GORM) - ports: 8080
- db: Postgres - ports: 5432

## ローカル起動（Dockerがある場合）
1. docker-compose build
2. docker-compose up

## Notes
- Backend depends on golang modules: echo, gorm, postgres driver.
  Add these to backend/go.mod or run `go get` locally if building outside Docker.
- Frontend uses Next.js; `npm install` will install dependencies in container.
