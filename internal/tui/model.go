// Package tui provides the interactive SSH portfolio interface.
package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	introFrames = 24
	introDelay  = 45 * time.Millisecond
)

type (
	introTickMsg     time.Time
	introCompleteMsg struct{}
)

// Model is the Bubble Tea model for the portfolio.
type Model struct {
	frame         int
	showPortfolio bool
	width         int
	height        int
}

// New returns a new portfolio model.
func New() Model {
	return Model{}
}

// Init begins the startup animation.
func (m Model) Init() tea.Cmd {
	return nextIntroFrame()
}

// Update handles terminal events and animation ticks.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case introTickMsg:
		if m.frame < introFrames {
			m.frame++
			return m, nextIntroFrame()
		}
		return m, tea.Tick(900*time.Millisecond, func(time.Time) tea.Msg {
			return introCompleteMsg{}
		})

	case introCompleteMsg:
		m.showPortfolio = true
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter", " ":
			m.showPortfolio = true
			return m, nil
		}
	}

	return m, nil
}

func nextIntroFrame() tea.Cmd {
	return tea.Tick(introDelay, func(t time.Time) tea.Msg {
		return introTickMsg(t)
	})
}
