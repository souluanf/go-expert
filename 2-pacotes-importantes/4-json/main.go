package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero  int     `json:"numero"`
	Agencia int     `json:"agencia"`
	Saldo   float64 `json:"-"` //Omitindo informação
}

func main() {
	conta := Conta{1, 2, 100.0}
	contaJson, err := json.Marshal(conta) //Guardo o valor para mim
	if err != nil {
		println(err)
	}
	println(string(contaJson))

	err = json.NewEncoder(os.Stdout).Encode(conta) //Pego o valor, serializo e entrega para alguém
	if err != nil {
		println(err)
	}

	jsonPuro := []byte(`{"numero":1,"agencia":2,"s":100}`)
	var conta2 Conta
	err = json.Unmarshal(jsonPuro, &conta2)
	if err != nil {
		println(err)
	}
	fmt.Println(conta2)
}
