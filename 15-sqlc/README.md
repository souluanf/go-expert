# Migrate

## Create a migration
```bash
migrate create -ext=sql -dir=sql/migrations -seq init
```

## Run migrations 
1. UP
```bash
migrate -path sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up
```
2. DOWN
```bash
migrate -path sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down
```

## Generate sqlc
```bash
sqlc generate
```