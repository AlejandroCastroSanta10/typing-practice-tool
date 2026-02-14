package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"typing-practice-tool/internal/lesson"
	"typing-practice-tool/internal/typing"
)

type tutorialState int

const (
	tutorialSelecting tutorialState = iota
	tutorialTyping
	tutorialComplete
)

type TutorialModel struct {
	lessons  []lesson.Lesson
	cursor   int
	state    tutorialState
	current  *lesson.Lesson
	tracker  *typing.Tracker
	width    int
	height   int
}

func NewTutorialModel() TutorialModel {
	return TutorialModel{
		lessons: lesson.AllLessons(),
	}
}

type SwitchToMenu struct{}

func (m TutorialModel) Init() tea.Cmd {
	return nil
}

func (m TutorialModel) Update(msg tea.Msg) (TutorialModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch m.state {
		case tutorialSelecting:
			return m.updateSelecting(msg)
		case tutorialTyping:
			return m.updateTyping(msg)
		case tutorialComplete:
			return m.updateComplete(msg)
		}
	}
	return m, nil
}

func (m TutorialModel) updateSelecting(msg tea.KeyMsg) (TutorialModel, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(m.lessons)-1 {
			m.cursor++
		}
	case "enter":
		m.current = &m.lessons[m.cursor]
		m.tracker = typing.NewTracker(m.current.GenerateExercise(60))
		m.state = tutorialTyping
	case "esc":
		return m, func() tea.Msg { return SwitchToMenu{} }
	}
	return m, nil
}

func (m TutorialModel) updateTyping(msg tea.KeyMsg) (TutorialModel, tea.Cmd) {
	if msg.String() == "esc" {
		m.state = tutorialSelecting
		return m, nil
	}
	runes := []rune(msg.String())
	if len(runes) == 1 {
		done := m.tracker.Type(runes[0])
		if done {
			m.state = tutorialComplete
		}
	}
	return m, nil
}

func (m TutorialModel) updateComplete(msg tea.KeyMsg) (TutorialModel, tea.Cmd) {
	switch msg.String() {
	case "enter":
		m.tracker.Reset()
		m.tracker = typing.NewTracker(m.current.GenerateExercise(60))
		m.state = tutorialTyping
	case "esc":
		m.state = tutorialSelecting
	}
	return m, nil
}

func (m TutorialModel) View() string {
	switch m.state {
	case tutorialSelecting:
		return m.viewSelecting()
	case tutorialTyping:
		return m.viewTyping()
	case tutorialComplete:
		return m.viewComplete()
	}
	return ""
}

func (m TutorialModel) viewSelecting() string {
	title := titleStyle.Render("Tutorial - Select a Lesson")

	var list string
	lastGroup := ""
	for i, l := range m.lessons {
		if l.Group != lastGroup {
			if lastGroup != "" {
				list += "\n"
			}
			list += subtitleStyle.Render(l.Group) + "\n"
			lastGroup = l.Group
		}
		cursor := "  "
		style := normalStyle
		if i == m.cursor {
			cursor = "> "
			style = selectedStyle
		}
		list += fmt.Sprintf("%s%s\n", cursor, style.Render(l.Name))
	}

	help := dimStyle.Render("↑/↓ navigate • enter select • esc back")

	content := lipgloss.JoinVertical(lipgloss.Left, title, "", list, help)
	box := boxStyle.Render(content)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, box)
}

func (m TutorialModel) viewTyping() string {
	title := titleStyle.Render(fmt.Sprintf("Lesson: %s", m.current.Name))

	// Finger hint
	hint := ""
	if m.tracker.Pos < len(m.tracker.Text) {
		ch := m.tracker.Text[m.tracker.Pos]
		finger := lesson.FingerName(ch)
		if finger != "" {
			hint = hintStyle.Render(fmt.Sprintf("Key '%s' → %s", string(ch), finger))
		}
	}

	textWidth := m.width - 8
	if textWidth < 40 {
		textWidth = 40
	}
	if textWidth > 80 {
		textWidth = 80
	}
	text := m.tracker.StyledText(textWidth)

	accuracy := fmt.Sprintf("Accuracy: %.0f%%", m.tracker.Accuracy())
	stats := statStyle.Render(accuracy)

	help := dimStyle.Render("esc back to lessons")

	content := lipgloss.JoinVertical(lipgloss.Left, title, hint, "", text, "", stats, "", help)
	box := boxStyle.Render(content)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, box)
}

func (m TutorialModel) viewComplete() string {
	title := titleStyle.Render("Lesson Complete!")

	accuracy := fmt.Sprintf("Accuracy: %.0f%%", m.tracker.Accuracy())
	stats := statStyle.Render(accuracy)

	help := dimStyle.Render("enter retry • esc back to lessons")

	content := lipgloss.JoinVertical(lipgloss.Left, title, "", stats, "", help)
	box := boxStyle.Render(content)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, box)
}
