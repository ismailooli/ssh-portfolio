package main

import (
	"fmt"
	"log"

	ssh "github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
)

func main() {
	server, err := wish.NewServer(
		wish.WithAddress(":2222"),
		wish.WithMiddleware(
			func(next ssh.Handler) ssh.Handler {
				return func(session ssh.Session) {
					fmt.Fprintln(session, "Welcome to my SSH portfolio!")
				}
			},
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("SSH server starting on port 2222...")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
