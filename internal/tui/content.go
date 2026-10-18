package tui

var bootLines = []string{
	"[ OK ] Establishing secure connection",
	"[ OK ] Verifying visitor identity",
	"[ OK ] Testing to see if anyone reading 19:15:38 ",
	"[ OK ] Verifying LinkedIn Warrior Status",
	"[ 404 ] Yo this guy capping on his linkedin",
	"[ OK ] Ahh I'll let it pass this one time",
	"[ OK ] Calculating storage needed for visit",
	"[ OK ] Ordering takeout from the local halal burger shop",
	"[ OK ] Picking up takeout from front door ",
	"[ OK ] Loading portfolio data",
	"[ OK ] System ready",
}

var menuItems = []string{
	"about me",
	"software projects",
	"film projects",
	"contact",
}

type Project struct {
	Title       string
	Description string
	Tags        []string
	URL         string
}

var softwareProjects = []Project{
	{
		Title:       "SSH Portfolio",
		Description: "A TUI portfolio served over SSH.",
		Tags:        []string{"Go", "Bubble Tea", "Lipgloss"},
		URL:         "https://github.com/ismailooli/ssh-portfolio",
	},
	{
		Title:       "slack-resizer",
		Description: "A tool to help resize images and gifs to meet the slack requirements + be able to preview them",
		Tags:        []string{"React", "Tailwind"},
		URL:         "https://slackresizer.com/",
	},
}

var filmProjects = []Project{
	{
		Title:       "The Last Shura Memeber",
		Description: "made a short film (full of inside jokes that you wont understand) to win the MSA film contest",
		Tags:        []string{"Director", "Editor"},
		URL:         "https://www.youtube.com/watch?v=vhSQCmEzz08&feature=youtu.be",
	},
}
