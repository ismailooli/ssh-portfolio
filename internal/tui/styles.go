package tui

import "github.com/charmbracelet/lipgloss"

const (
	menuWidth    = 22
	contentWidth = 80
	panelHeight  = 24
)

var (
	borderColor     = lipgloss.Color("#5A6A73")
	logoStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#35D0D6"))
	dimStyle        = lipgloss.NewStyle().Foreground(borderColor)
	menuActiveStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#6FE7E1")).
			Foreground(lipgloss.Color("#1F2435")).
			Width(menuWidth).
			Padding(0, 1)
	menuItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#8C96A8")).
			Width(menuWidth).
			Padding(0, 1)
	leftPanelStyle = lipgloss.NewStyle().
			BorderRight(true).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(borderColor).
			Padding(1, 2).
			Height(panelHeight)
	rightPanelStyle = lipgloss.NewStyle().
			BorderLeft(true).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(borderColor).
			Padding(1, 2).
			Height(panelHeight)
)
