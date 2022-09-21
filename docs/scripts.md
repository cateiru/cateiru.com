# Scripts

## Connect to Local SQL

```bash
./scripts/sql.sh
```

## Migration Local SQL

```bash
go run ./scripts/main.go migration [mode]
```

- mode: Config mode. `test`, `local` or `prod`
  - If no mode is set, it will migrate the `test` table by default.

## Export SQL Schema

```bash
go run ./scripts/main.go export
```
