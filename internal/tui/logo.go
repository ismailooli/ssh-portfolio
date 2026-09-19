package tui

import (
	_ "embed"
	"strings"
)

//go:embed logo.txt
var logoRaw string

var logo = strings.Split(strings.TrimRight(logoRaw, "\n"), "\n")

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
