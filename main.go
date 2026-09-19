package main

import (
	"log"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	ssh "github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	wishbubbletea "github.com/charmbracelet/wish/bubbletea"
)

const (
	introFrames = 24
	introDelay  = 45 * time.Millisecond
)

var (
	logoStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#35D0D6"))
	dimStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#5A6A73"))
)

// This is deliberately plain text art: every character can be changed to make
// the mark feel like your own.
var logo = []string{
	"                 · · · · · · · · · · ·",
	"           · · ·                       · · ·",
	"       · ·       · · · · · · · · · ·       · ·",
	"     · ·      · ·                   · ·      · ·",
	"    · ·     · ·                         · ·     · ·",
	"    · ·    · ·      I S M A I L          · ·    · ·",
	"    · ·     · ·                         · ·     · ·",
	"     · ·      · ·     PORTFOLIO      · ·      · ·",
	"       · ·       · · · · · · · · · ·       · ·",
	"           · · ·                       · · ·",
	"                 · · · · · · · · · · ·",
}

type introTickMsg time.Time
type introCompleteMsg struct{}

type model struct {
	frame         int
	showPortfolio bool
	width         int
	height        int
}

func (m model) Init() tea.Cmd {
	return nextIntroFrame()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case introTickMsg:
		if m.frame < introFrames {
			m.frame++
			return m, nextIntroFrame()
		}
		return m, tea.Tick(900*time.Millisecond, func(time.Time) tea.Msg {
			return introCompleteMsg{}
		})

	case introCompleteMsg:
		m.showPortfolio = true
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter", " ":
			// Let someone skip the intro while you are developing it.
			m.showPortfolio = true
			return m, nil
		}
	}

	return m, nil
}

func (m model) View() string {
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

func (m model) introView() string {
	art := logoStyle.Render(revealLogo(m.frame))
	progress := strings.Repeat("█", m.frame/2) + strings.Repeat("░", introFrames/2-m.frame/2)

	return art + "\n\n" + dimStyle.Render("  establishing secure connection...") +
		"\n" + logoStyle.Render("  ["+progress+"]") +
		"\n\n" + dimStyle.Render("  press enter to skip")
}

// revealLogo exposes a little more of the text art for every animation frame.
func revealLogo(frame int) string {
	var total int
	for _, line := range logo {
		for _, char := range line {
			if char != ' ' {
				total++
			}
		}
	}

	visible := total * frame / introFrames
	var output strings.Builder
	for lineIndex, line := range logo {
		for _, char := range line {
			if char == ' ' {
				output.WriteRune(' ')
				continue
			}
			if visible > 0 {
				output.WriteRune(char)
				visible--
			} else {
				output.WriteRune(' ')
			}
		}
		if lineIndex < len(logo)-1 {
			output.WriteByte('\n')
		}
	}

	return output.String()
}

func (m model) place(content string) string {
	if m.width == 0 || m.height == 0 {
		return content
	}
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)
}

func nextIntroFrame() tea.Cmd {
	return tea.Tick(introDelay, func(t time.Time) tea.Msg {
		return introTickMsg(t)
	})
}

func teaHandler(session ssh.Session) (tea.Model, []tea.ProgramOption) {
	return model{}, []tea.ProgramOption{
		tea.WithAltScreen(),
	}
}

func main() {
	server, err := wish.NewServer(
		wish.WithAddress(":2222"),
		wish.WithHostKeyPath(".ssh-portfolio-host-key"),
		wish.WithMiddleware(
			wishbubbletea.Middleware(teaHandler),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("SSH server starting on port 2222...")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
