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
	clockTickMsg     time.Time
	introCompleteMsg struct{}
)

type screen int

const (
	bootScreen screen = iota
	logoScreen
	portfolioScreen
)

type Model struct {
	screen       screen
	logLine      int
	frame        int
	width        int
	height       int
	selected     int
	currentTime  time.Time
	selectedCard int
	expandedCard int
}

func New() Model {
	return Model{
		currentTime:  time.Now(),
		expandedCard: -1,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(nextBootLine(), nextClockTick())
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

	case clockTickMsg:
		m.currentTime = time.Time(msg)
		return m, nextClockTick()

	case introCompleteMsg:
		m.screen = portfolioScreen
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "enter", " ":
			// Enter skips the boot/logo screen.
			if m.screen != portfolioScreen {
				m.screen = portfolioScreen
				return m, nil
			}

			// Enter expands the focused software or film project card.
			projects := m.activeProjects()
			if len(projects) > 0 && m.expandedCard == -1 {
				m.expandedCard = m.selectedCard
				return m, nil
			}

		case "esc":
			// Escape only closes an expanded project.
			if m.expandedCard != -1 {
				m.expandedCard = -1
				return m, nil
			}

		case "tab":
			projects := m.activeProjects()
			if len(projects) > 0 && m.expandedCard == -1 {
				m.selectedCard = (m.selectedCard + 1) % len(projects)
				return m, nil
			}

		case "shift+tab":
			projects := m.activeProjects()
			if len(projects) > 0 && m.expandedCard == -1 {
				m.selectedCard = (m.selectedCard - 1 + len(projects)) % len(projects)
				return m, nil
			}

		case "down", "j":
			if m.screen == portfolioScreen &&
				m.expandedCard == -1 &&
				m.selected < len(menuItems)-1 {
				m.selected++
				m.selectedCard = 0
			}

		case "up", "k":
			if m.screen == portfolioScreen &&
				m.expandedCard == -1 &&
				m.selected > 0 {
				m.selected--
				m.selectedCard = 0
			}
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

func nextClockTick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return clockTickMsg(t)
	})
}
