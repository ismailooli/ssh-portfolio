package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	logoStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#35D0D6"))
	dimStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#5A6A73"))
	menuActiveStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#6FE7E1")).
			Foreground(lipgloss.Color("#1F2435")).
			Width(22).
			Padding(0, 1)

	menuItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#8C96A8")).
			Width(22).
			Padding(0, 1)
)

// View renders the boot and logo animations, then a blank portfolio canvas.
func (m Model) View() string {
	switch m.screen {
	case bootScreen:
		return m.placeBottomLeft(m.bootView())

	case logoScreen:
		return m.place(m.introView())

	case portfolioScreen:
		return m.place(m.menuView())
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

func (m Model) menuView() string {
	items := []string{"about me", "projects"}

	var menuRows []string
	for index, item := range items {
		if index == m.selected {
			menuRows = append(menuRows, menuActiveStyle.Render(item))
		} else {
			menuRows = append(menuRows, menuItemStyle.Render(item))
		}
	}

	menu := strings.Join(menuRows, "\n")

	content := ""
	switch m.selected {
	case 0:
		content = `about me

		I’m Ismail—a developer and filmmaker.

		This SSH portfolio is a small terminal
		interface built with Go and Bubble Tea.`

	case 1:
		content = `projects

		SSH Portfolio
		A terminal-based portfolio served over SSH.`
	}

	const (
		menuWidth    = 22
		contentWidth = 100
		panelHeight  = 25
	)

	borderColor := lipgloss.Color("#5A6A73")

	rightPanelStyle := lipgloss.NewStyle().
		BorderLeft(true).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(borderColor).
		Padding(1, 2).
		Height(panelHeight)

	leftPanelStyle := lipgloss.NewStyle().
		BorderRight(true).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(borderColor).
		Padding(1, 2).
		Height(panelHeight)

	menuPanel := leftPanelStyle.Width(menuWidth).Render(menu)
	contentPanel := rightPanelStyle.Width(contentWidth).Render(content)

	body := lipgloss.JoinHorizontal(
		lipgloss.Top,
		menuPanel,
		contentPanel,
	)

	headerStyle := lipgloss.NewStyle().
		Width(lipgloss.Width(body)).
		BorderBottom(true).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(borderColor)

	footerStyle := lipgloss.NewStyle().
		Width(lipgloss.Width(body)).
		BorderTop(true).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(borderColor).
		Foreground(lipgloss.Color("#5A6A73"))

	header := headerStyle.Render(
		"ismail mohammad                              [portfolio]",
	)

	footer := footerStyle.Render(
		"↑/k and ↓/j to navigate • q to quit",
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		body,
		footer,
	)
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
