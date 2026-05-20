#Knowledge-Game

# Migrations:
```bash
go install github.com/rubenv/sql-migrate/...@latest
```

To migrate up: 
```bash
sql-migrate up -env=production
```

To migrate down: 
```bash
sql-migrate down -env=production
```
