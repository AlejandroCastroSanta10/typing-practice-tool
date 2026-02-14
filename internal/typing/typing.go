package typing

import (
	"time"

	"github.com/charmbracelet/lipgloss"
)

// Result tracks per-character correctness.
type Result int

const (
	Pending Result = iota
	Correct
	Incorrect
)

// Tracker handles character-by-character input matching and statistics.
type Tracker struct {
	Text    []rune
	Results []Result
	Pos     int
	started bool
	start   time.Time
}

// NewTracker creates a tracker for the given exercise text.
func NewTracker(text string) *Tracker {
	runes := []rune(text)
	return &Tracker{
		Text:    runes,
		Results: make([]Result, len(runes)),
	}
}

// Type processes a single keystroke and returns whether the exercise is complete.
func (t *Tracker) Type(ch rune) bool {
	if t.Pos >= len(t.Text) {
		return true
	}
	if !t.started {
		t.started = true
		t.start = time.Now()
	}
	if ch == t.Text[t.Pos] {
		t.Results[t.Pos] = Correct
	} else {
		t.Results[t.Pos] = Incorrect
	}
	t.Pos++
	return t.Pos >= len(t.Text)
}

// CorrectCount returns the number of correctly typed characters.
func (t *Tracker) CorrectCount() int {
	count := 0
	for _, r := range t.Results {
		if r == Correct {
			count++
		}
	}
	return count
}

// TypedCount returns the total number of characters typed so far.
func (t *Tracker) TypedCount() int {
	return t.Pos
}

// WPM returns the current words-per-minute (gross WPM).
func (t *Tracker) WPM() float64 {
	if !t.started || t.Pos == 0 {
		return 0
	}
	elapsed := time.Since(t.start).Minutes()
	if elapsed < 0.001 {
		return 0
	}
	return (float64(t.CorrectCount()) / 5.0) / elapsed
}

// Accuracy returns the current accuracy as a percentage (0-100).
func (t *Tracker) Accuracy() float64 {
	if t.Pos == 0 {
		return 100
	}
	return float64(t.CorrectCount()) / float64(t.Pos) * 100
}

// Done returns true if all characters have been typed.
func (t *Tracker) Done() bool {
	return t.Pos >= len(t.Text)
}

// Reset restarts the tracker with the same text.
func (t *Tracker) Reset() {
	t.Pos = 0
	t.Results = make([]Result, len(t.Text))
	t.started = false
}

// StyledText renders the exercise text with colors: green for correct,
// red for incorrect, dimmed for untyped. The cursor position is underlined.
func (t *Tracker) StyledText(width int) string {
	correctStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
	incorrectStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("1"))
	dimStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	cursorStyle := lipgloss.NewStyle().Underline(true)

	var out string
	col := 0
	for i, ch := range t.Text {
		s := string(ch)
		if ch == '\n' {
			out += "\n"
			col = 0
			continue
		}
		if width > 0 && col >= width {
			out += "\n"
			col = 0
		}
		switch {
		case i < t.Pos:
			if t.Results[i] == Correct {
				out += correctStyle.Render(s)
			} else {
				out += incorrectStyle.Render(s)
			}
		case i == t.Pos:
			out += cursorStyle.Render(s)
		default:
			out += dimStyle.Render(s)
		}
		col++
	}
	return out
}
