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

To generate proto files:
```bash
protoc -I=rpcmessage --go_out=rpcmessage/golang/presence --go_opt=paths=source_relative --go-grpc_out=rpcmessage/golang/presence --go-grpc_opt=paths=source_relative rpcmessage/presence.proto
```


