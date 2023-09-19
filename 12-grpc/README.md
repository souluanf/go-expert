# gRPC

## Generate gRPC code

```shell
$ protoc --go_out=. --go-grpc_out=. proto/course_category.proto
```

## create table using sqlite3
```shell
$ sqlite3 db.sqlite3 
```

```sql
create table categories
(
    id          string,
    name        string,
    description string
);
```

## Run server

```shell
$ go run server/main.go
```

## Connect to server using evans

```shell
evans -r repl
```

```shell
$ package pb
```
```shell
$ service CategoryService
```
```shell
$ call CreateCategory
```
```shell
