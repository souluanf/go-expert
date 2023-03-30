#### Execurando todos os testes
```shell
go test . 
go test -v //verboso
```

#### verificando cobertura
```shell
go test -cover //mostra a cobertura sem arquivo
go test -coverprofile=coverage.out //gera um arquivo de cobertura
```

#### gera html com a cobertura
```shell
go tool cover -html=coverage.out
```

#### rodando benchmark
```shell
go test -bench=.
```

#### rodando benchmark sem os testes
```shell
go test -bench=. -run=^#
```

#### rodando benchmark sem os testes 10 vezes
```shell
go test -bench=. -run=^# -count=10
```