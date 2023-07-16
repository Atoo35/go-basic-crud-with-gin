package main

import (
	"fmt"
	"log"

	"github.com/Atoo35/basic-crud/configurations"
)

func main() {
	config, err := configurations.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(config.TestVar)
}
