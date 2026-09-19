// Package server configures and runs the portfolio SSH server.
package server

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	ssh "github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	wishbubbletea "github.com/charmbracelet/wish/bubbletea"

	"ssh-portfolio/internal/tui"
)

const (
	defaultAddress     = ":2222"
	defaultHostKeyPath = ".ssh-portfolio-host-key"
)

type Config struct {
	Address     string
	HostKeyPath string
}

func Run(config Config) error {
	address := config.Address
	if address == "" {
		address = defaultAddress
	}

	hostKeyPath := config.HostKeyPath
	if hostKeyPath == "" {
		hostKeyPath = defaultHostKeyPath
	}

	server, err := wish.NewServer(
		wish.WithAddress(address),
		wish.WithHostKeyPath(hostKeyPath),
		wish.WithMiddleware(wishbubbletea.Middleware(teaHandler)),
	)
	if err != nil {
		return err
	}

	log.Printf("SSH server starting on %s...", address)
	return server.ListenAndServe()
}

func teaHandler(ssh.Session) (tea.Model, []tea.ProgramOption) {
	return tui.New(), []tea.ProgramOption{tea.WithAltScreen()}
}
