package main

import (
	"github.com/felipefabricio/wonder-food/configs"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
}
