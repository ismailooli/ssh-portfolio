package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	ssh "github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	wishbubbletea "github.com/charmbracelet/wish/bubbletea"
)

type model struct{}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	return `
    Welcome to my SSH portfolio

    Press q to quit.
  `
}

func teaHandler(session ssh.Session) (tea.Model, []tea.ProgramOption) {
	return model{}, []tea.ProgramOption{
		tea.WithAltScreen(),
	}
}

func main() {
	server, err := wish.NewServer(
		wish.WithAddress(":2222"),
		wish.WithHostKeyPath(".ssh-portfolio-host-key"),
		wish.WithMiddleware(
			wishbubbletea.Middleware(teaHandler),
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
