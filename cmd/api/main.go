package main

import (
	"fmt"

	"github.com/dougmendes/advg/cmd/api/config"
)

func main() {
	fmt.Println("initializing...")
	server := config.NewServer()
	server.Run()
}
