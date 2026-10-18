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

var (
	projectCardStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(borderColor).
				Padding(1, 2).
				MarginBottom(1).
				Width(72)

	projectCardFocusedStyle = projectCardStyle.
				BorderForeground(lipgloss.Color("#6FE7E1"))

	expandedCardStyle = lipgloss.NewStyle().
				Border(lipgloss.DoubleBorder()).
				BorderForeground(lipgloss.Color("#6FE7E1")).
				Padding(2, 3).
				Width(72)

	projectTitleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#E8EDF4")).
				Bold(true)

	tagStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#3B465A")).
			Foreground(lipgloss.Color("#E8EDF4")).
			Padding(0, 1).
			MarginRight(1)

	linkStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6FE7E1")).
			Underline(true)
)
