package main

import (
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env/v6"
)

type config struct {
	CommandMode string `env:"COMMAND_MODE"`
}

func main() {
	args := os.Args
	fmt.Println(len(args))
	for _, a := range args {
		println(a)
	}
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)
}
