package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	logoStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#35D0D6"))
	dimStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#5A6A73"))
)

// View renders either the connection animation or the portfolio menu.
func (m Model) View() string {
	switch m.screen {
	case bootScreen:
		return m.placeBottomLeft(m.bootView())

	case logoScreen:
		return m.place(m.introView())

	case portfolioScreen:
		return m.place(`
    Welcome to my SSH portfolio

    [1] About me
    [2] Projects
    [3] Contact

    Press q to quit.
  `)
	}

	return ""
}

func (m Model) introView() string {
	art := logoStyle.Render(revealLogo(m.frame))

	return art + "\n\n" + dimStyle.Render(" giving you access to the secret sauce...")
}

func (m Model) bootView() string {
	return logoStyle.Render(strings.Join(bootLines[:m.logLine], "\n"))
}

func (m Model) place(content string) string {
	if m.width == 0 || m.height == 0 {
		return content
	}
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)
}

func (m Model) placeBottomLeft(content string) string {
	if m.width == 0 || m.height == 0 {
		return content
	}

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Left,
		lipgloss.Bottom,
		content,
	)
}
