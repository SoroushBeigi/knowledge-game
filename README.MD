# Knowledge-Game

# Migrations:
```bash
go install github.com/rubenv/sql-migrate/...@latest
```

To migrate up: 
```bash
sql-migrate up -env=production -config=repository/mysql/dbconfig.yml
```

To migrate down: 
```bash
sql-migrate down -env=production -config=repository/mysql/dbconfig.yml
```

To see migration status: 
```bash
sql-migrate status -env=production -config=repository/mysql/dbconfig.yml
```


