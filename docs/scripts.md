# Scripts

## Connect to Local SQL

```bash
./scripts/sql.sh
```

## Migration Local SQL

```bash
go run ./scripts/main.go migration [OPTIONS]
```

- mode: Config mode. `test`, `local` or `prod`
  - If no mode is set, it will migrate the `local` table by default.
- env
  - https://github.com/cateiru/sql env file path.

### Connect CloudSQL

Use https://github.com/cateiru/sql

```bash
cd sql
# start gce-proxy
docker-compose up -d

cd ../cateiru.com
go run ./scripts/main.go migration -m prod -e ../sql/.env
# Export migration_diff.sql file.
```

## Export SQL Schema

```bash
go run ./scripts/main.go export
# Export schema.sql
```
