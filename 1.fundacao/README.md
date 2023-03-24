# Runtime

```mermaid
graph LR
A[Código Go Runtime] --> B[Seu código] --> C[Binário]
```

# Mostrar arquiteturas disponiveis
```shell
go tool dist list
```
# Build para sistema especifico
```shell
GOOS=linux GOARCH=amd64 go build
GOOS=windows GOARCH=amd64 go build
GOOS=darwin GOARCH=arm64 go build
```

# Build nomeado
```shell
GOOS=darwin GOARCH=arm64 go build -o nome_do_arquivo
```