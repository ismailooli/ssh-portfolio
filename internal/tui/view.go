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
	if m.showPortfolio {
		return m.place(`
  Welcome to my SSH portfolio

  [1] About me
  [2] Projects
  [3] Contact

  Press q to quit.
`)
	}

	return m.place(m.introView())
}

func (m Model) introView() string {
	art := logoStyle.Render(revealLogo(m.frame))
	progress := strings.Repeat("█", m.frame/2) + strings.Repeat("░", introFrames/2-m.frame/2)

	return art + "\n\n" + dimStyle.Render("  establishing secure connection...") +
		"\n" + logoStyle.Render("  ["+progress+"]")
}

func (m Model) place(content string) string {
	if m.width == 0 || m.height == 0 {
		return content
	}
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)
}
