# wolfort-games backend

### GraphQL スキーマ定義変更

1. ../graphql/schema.graphqls を変更
1. `go get github.com/99designs/gqlgen@v0.17.40;go run github.com/99designs/gqlgen generate`

### 環境変数

.env を用意

例

```
DB_HOST=localhost
DB_NAME=gamesdb
DB_USER=games
DB_PASS=password
OAUTH_ISSUERURI={oauth issuer uri}
OAUTH_AUDIENCE={oauth audience (comma separated)}
```

### サーバー起動

```
go run main.go
```

### library

- GraphQL: gqlgen
- ORM: gorm
- Auth: Auth0 JWT

### todo

- over fetch を防ぐ
  - see: https://tech.layerx.co.jp/entry/2021/10/22/171242
