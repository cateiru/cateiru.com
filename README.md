# https://cateiru.com

## Edit SQL Schema

[edit schema using ent](./ent/schema)

## Environments

```env
DB_USER=database user name
DB_PASSWORD=database password
MAILGUN_APIKEY=mailgun api key
SSO_CLIENT_ID=cateiru sso client id
SSO_TOKEN_SECRET=cateiru sso token secret
NEXT_PUBLIC_API_DOMAIN=api domain
```

## Setup Local

```bash
# start mysql
docker compose up -d

# download deps
go mod download
yarn

# migration
go run ./script/main.go migrate --mode local
go run ./script/main.go migrate --mode test

# start
go run .
yarn dev

# access to http://localhost:3000
```

## Scripts

[see here](./docs/scripts.md)

## LICENSE

[MIT](./LICENSE)
