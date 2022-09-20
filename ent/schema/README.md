# SQL SChema

Using [ent](https://entgo.io/)

## Add Schema

```bash
# add `User` and `Pet` table
go run -mod=mod entgo.io/ent/cmd/ent init User Pet
```

## Generate Scripts

```bash
go generate ./ent
```

## Export SQL

```bash
# start local mysql server
docker-compose up -d

# export
go run ./scripts/main.go export
```
