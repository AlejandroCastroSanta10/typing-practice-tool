package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"typing-practice-tool/internal/typing"
	"typing-practice-tool/internal/wordlist"
)

type practiceState int

const (
	practiceTyping practiceState = iota
	practiceSummary
)

type PracticeModel struct {
	tracker    *typing.Tracker
	state      practiceState
	width      int
	height     int
	totalChars int
	finalWPM   float64
	finalAcc   float64
}

func NewPracticeModel() PracticeModel {
	text := wordlist.Generate(200)
	return PracticeModel{
		tracker: typing.NewTracker(text),
	}
}

func (m PracticeModel) Init() tea.Cmd {
	return nil
}

func (m PracticeModel) Update(msg tea.Msg) (PracticeModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch m.state {
		case practiceTyping:
			return m.updateTyping(msg)
		case practiceSummary:
			return m.updateSummary(msg)
		}
	}
	return m, nil
}

func (m PracticeModel) updateTyping(msg tea.KeyMsg) (PracticeModel, tea.Cmd) {
	if msg.String() == "esc" {
		m.finalWPM = m.tracker.WPM()
		m.finalAcc = m.tracker.Accuracy()
		m.totalChars = m.tracker.TypedCount()
		m.state = practiceSummary
		return m, nil
	}
	runes := []rune(msg.String())
	if len(runes) == 1 {
		done := m.tracker.Type(runes[0])
		if done {
			// Continuous flow: generate new words and keep going
			m.totalChars += m.tracker.TypedCount()
			text := wordlist.Generate(200)
			m.tracker = typing.NewTracker(text)
		}
	}
	return m, nil
}

func (m PracticeModel) updateSummary(msg tea.KeyMsg) (PracticeModel, tea.Cmd) {
	// Any key returns to menu
	return m, func() tea.Msg { return SwitchToMenu{} }
}

func (m PracticeModel) View() string {
	switch m.state {
	case practiceTyping:
		return m.viewTyping()
	case practiceSummary:
		return m.viewSummary()
	}
	return ""
}

func (m PracticeModel) viewTyping() string {
	title := titleStyle.Render("Free Practice")

	textWidth := m.width - 8
	if textWidth < 40 {
		textWidth = 40
	}
	if textWidth > 80 {
		textWidth = 80
	}
	text := m.tracker.StyledText(textWidth)

	wpm := fmt.Sprintf("WPM: %.0f", m.tracker.WPM())
	acc := fmt.Sprintf("Accuracy: %.0f%%", m.tracker.Accuracy())
	stats := statStyle.Render(fmt.Sprintf("%s  |  %s", wpm, acc))

	help := dimStyle.Render("esc finish")

	content := lipgloss.JoinVertical(lipgloss.Left, title, "", text, "", stats, "", help)
	box := boxStyle.Render(content)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, box)
}

func (m PracticeModel) viewSummary() string {
	title := titleStyle.Render("Session Summary")

	wpm := fmt.Sprintf("WPM: %.0f", m.finalWPM)
	acc := fmt.Sprintf("Accuracy: %.0f%%", m.finalAcc)
	chars := fmt.Sprintf("Characters typed: %d", m.totalChars)

	stats := lipgloss.JoinVertical(lipgloss.Left,
		statStyle.Render(wpm),
		statStyle.Render(acc),
		statStyle.Render(chars),
	)

	help := dimStyle.Render("press any key to return to menu")

	content := lipgloss.JoinVertical(lipgloss.Left, title, "", stats, "", help)
	box := boxStyle.Render(content)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, box)
}
