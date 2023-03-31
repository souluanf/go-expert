package main

import (
	"fmt"
	"github.com/souluanf/go-expert/7-apis/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	fmt.Println(config.DBName)
	fmt.Println(config.TokenAuth)
}
