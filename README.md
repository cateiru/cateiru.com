# https://cateiru.com

## Edit SQL Schema

[edit schema using ent](./ent/schema)

## Environments

```env
DB_USER=[database user name]
DB_PASSWORD=[database password]
MAILGUN_APIKEY=[mailgun api key]
SSO_CLIENT_ID=[cateiru sso client id]
SSO_TOKEN_SECRET=[cateiru sso token secret]
NEXT_PUBLIC_API_DOMAIN=[api domain]
NEXT_PUBLIC_BACKEND_API_DOMAIN=[backend api domain]
```

## Setup Local

```bash
./docker-compose up -d

# migration
go run ./scripts/main.go migration --mode local
go run ./scripts/main.go migration --mode test

# access to http://localhost:3000
```

## Scripts

[see here](./docs/scripts.md)

## LICENSE

[MIT](./LICENSE)
