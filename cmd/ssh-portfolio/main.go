package main

import (
	"log"

	"ssh-portfolio/internal/server"
)

func main() {
	if err := server.Run(server.Config{}); err != nil {
		log.Fatal(err)
	}
}
