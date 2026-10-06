package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// View renders the boot and logo animations, then a blank portfolio canvas.
func (m Model) View() string {
	switch m.screen {
	case bootScreen:
		return m.placeBottomLeft(m.bootView())

	case logoScreen:
		return m.place(m.introView())

	case portfolioScreen:
		return m.place(m.portfolioView())
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
