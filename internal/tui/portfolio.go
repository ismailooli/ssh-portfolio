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

	footer := footerStyle.Render(
		"↑/↓ sections • tab choose project • enter open • esc back • q quit",
	)

	return lipgloss.JoinVertical(lipgloss.Left, header, body, footer)
}

func (m Model) portfolioContent() string {
	switch m.selected {
	case 0:
		return `
		about me

		Hi, my name is ismail -- i'm currently working as a swe at Relativity working on our FOIA platform but 
		in my free time, I like to mess around with terminal tools and watch crappy movies`
	case 1, 2:
		projects := m.activeProjects()

		if m.expandedCard != -1 {
			return m.expandedProject(projects[m.expandedCard])
		}

		return m.projectCards(projects)

	case 3:
		return `contact information goes here`
	default:
		return ""
	}
}

// Project Card list
func (m Model) activeProjects() []Project {
	switch m.selected {
	case 1:
		return softwareProjects
	case 2:
		return filmProjects
	default:
		return nil
	}
}

func (m Model) expandedProject(project Project) string {
	tags := make([]string, len(project.Tags))
	for index, tag := range project.Tags {
		tags[index] = tagStyle.Render(tag)
	}

	link := linkStyle.Render(project.URL)

	content := projectTitleStyle.Render(project.Title) +
		"\n\n" + project.Description +
		"\n\n" + strings.Join(tags, " ") +
		"\n\n" + link +
		"\n\n" + dimStyle.Render("esc to return to all projects")

	return expandedCardStyle.Render(content)
}

func (m Model) projectCards(projects []Project) string {
	var cards []string

	for index, project := range projects {
		style := projectCardStyle
		marker := "  "

		if index == m.selectedCard {
			style = projectCardFocusedStyle
			marker = "▶ "
		}

		tags := make([]string, len(project.Tags))
		for tagIndex, tag := range project.Tags {
			tags[tagIndex] = tagStyle.Render(tag)
		}

		card := marker + projectTitleStyle.Render(project.Title) +
			"\n\n" + dimStyle.Render(project.Description) +
			"\n\n" + strings.Join(tags, " ")

		cards = append(cards, style.Render(card))
	}

	return strings.Join(cards, "\n")
}
