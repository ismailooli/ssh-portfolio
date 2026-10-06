package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) portfolioView() string {
	var menuRows []string
	for index, item := range menuItems {
		if index == m.selected {
			menuRows = append(menuRows, menuActiveStyle.Render(item))
		} else {
			menuRows = append(menuRows, menuItemStyle.Render(item))
		}
	}

	menu := strings.Join(menuRows, "\n")
	content := m.portfolioContent()

	menuPanel := leftPanelStyle.Width(menuWidth).Render(menu)
	contentPanel := rightPanelStyle.Width(contentWidth).Render(content)
	body := lipgloss.JoinHorizontal(lipgloss.Top, menuPanel, contentPanel)

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
		Foreground(borderColor)

	left := "ismail mohammad"
	right := m.currentTime.Local().Format("3:04:05 PM")
	header := headerStyle.Render(
		left + lipgloss.PlaceHorizontal(
			max(0, lipgloss.Width(body)-lipgloss.Width(left)),
			lipgloss.Right,
			right,
		),
	)
	footer := footerStyle.Render("↑/k and ↓/j to navigate • q to quit")

	return lipgloss.JoinVertical(lipgloss.Left, header, body, footer)
}

func (m Model) portfolioContent() string {
	switch m.selected {
	case 0:
		return `
		about me

		Hi, my name is ismail -- i'm currently working as a swe at Relativity working on our FOIA platform but 
		in my free time, I like to mess around with terminal tools and watch crappy movies`
	case 1:
		return `
		projects

		SSH Portfolio
		A terminal-based portfolio served over SSH.`
	case 2, 3:
		return ` insert placeholder here`
	default:
		return ""
	}
}
