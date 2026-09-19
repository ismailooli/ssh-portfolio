package tui

import (
	"strings"
	"testing"
)

func TestRevealLogo(t *testing.T) {
	if got := revealLogo(0); stringsCountNonSpace(got) != 0 {
		t.Fatalf("revealLogo(0) shows %d visible characters, want 0", stringsCountNonSpace(got))
	}

	if got, want := revealLogo(introFrames), strings.Join(logo, "\n"); got != want {
		t.Fatal("revealLogo at the final frame does not show the full logo")
	}
}

func stringsCountNonSpace(value string) int {
	count := 0
	for _, char := range value {
		if char != ' ' && char != '\n' {
			count++
		}
	}
	return count
}
