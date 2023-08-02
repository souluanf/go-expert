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

#### Rodando benchmark sem os testes

```shell
go test -bench=. -run=^#
```

#### Rodando benchmark sem os testes 10 vezes

```shell
go test -bench=. -run=^# -count=10
```

#### rodar fuzzing

```shell
go test -fuzz=.
```

#### rodar fuzzing com tempo

```shell
go test -fuzz=. -fuzztime=10s
```