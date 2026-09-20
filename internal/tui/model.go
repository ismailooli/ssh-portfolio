// Package tui provides the interactive SSH portfolio interface.
package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	introFrames = 24
	introDelay  = 45 * time.Millisecond
	bootDelay   = 267 * time.Millisecond
)

type (
	bootTickMsg      time.Time
	introTickMsg     time.Time
	introCompleteMsg struct{}
)

type screen int

const (
	bootScreen screen = iota
	logoScreen
	portfolioScreen
)

type Model struct {
	screen  screen
	logLine int
	frame   int
	width   int
	height  int
}

var bootLines = []string{
	"[ OK ] Establishing secure connection",
	"[ OK ] Verifying visitor identity",
	"[ OK ] Testing to see if anyone reading 19:15:38 ",
	"[ OK ] Verifying LinkedIn Warrior Status",
	"[ 404 ] Yo this guy capping on his linkedin",
	"[ OK ] Ahh I'll let it pass this one time",
	"[ OK ] Calculating storage needed for visit",
	"[ OK ] Ordering takeout from the local halal burger shop",
	"[ OK ] Picking up takeout from front door ",
	"[ OK ] Loading portfolio data",
	"[ OK ] System ready",
}

func New() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	return nextBootLine()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case bootTickMsg:
		if m.logLine < len(bootLines) {
			m.logLine++
			return m, nextBootLine()
		}

		m.screen = logoScreen
		return m, nextIntroFrame()

	case introTickMsg:
		if m.frame < introFrames {
			m.frame++
			return m, nextIntroFrame()
		}
		return m, tea.Tick(900*time.Millisecond, func(time.Time) tea.Msg {
			return introCompleteMsg{}
		})

	case introCompleteMsg:
		m.screen = portfolioScreen
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter", " ":
			m.screen = portfolioScreen
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

func nextBootLine() tea.Cmd {
	return tea.Tick(bootDelay, func(t time.Time) tea.Msg {
		return bootTickMsg(t)
	})
}
