package main

import (
	"log"
	"os"

	"github.com/smiyaguchi/forger/internal/config"
)

func run() error {
	_, err := config.Load("./test/data/test.yml")
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Printf("error occured: %#v\n", err)
		os.Exit(1)
	}
}
