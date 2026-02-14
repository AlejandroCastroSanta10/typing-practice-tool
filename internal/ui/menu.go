package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type menuChoice int

const (
	menuTutorial menuChoice = iota
	menuPractice
	menuQuit
)

type MenuModel struct {
	cursor int
	width  int
	height int
}

func NewMenuModel() MenuModel {
	return MenuModel{}
}

type SwitchToTutorial struct{}
type SwitchToPractice struct{}

func (m MenuModel) Init() tea.Cmd {
	return nil
}

func (m MenuModel) Update(msg tea.Msg) (MenuModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < 2 {
				m.cursor++
			}
		case "enter":
			switch menuChoice(m.cursor) {
			case menuTutorial:
				return m, func() tea.Msg { return SwitchToTutorial{} }
			case menuPractice:
				return m, func() tea.Msg { return SwitchToPractice{} }
			case menuQuit:
				return m, tea.Quit
			}
		case "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m MenuModel) View() string {
	items := []string{"Tutorial", "Free Practice", "Quit"}

	title := titleStyle.Render("Typing Practice Tool")
	subtitle := dimStyle.Render("Learn to type correctly, one key at a time")

	var options string
	for i, item := range items {
		cursor := "  "
		style := normalStyle
		if i == m.cursor {
			cursor = "> "
			style = selectedStyle
		}
		options += fmt.Sprintf("%s%s\n", cursor, style.Render(item))
	}

	help := dimStyle.Render("↑/↓ navigate • enter select • q quit")

	content := lipgloss.JoinVertical(lipgloss.Left,
		title,
		subtitle,
		"",
		options,
		help,
	)

	box := boxStyle.Render(content)

	return lipgloss.Place(m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		box,
	)
}
