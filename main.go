package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"typing-practice-tool/internal/ui"
)

type screen int

const (
	screenMenu screen = iota
	screenTutorial
	screenPractice
)

type model struct {
	screen   screen
	menu     ui.MenuModel
	tutorial ui.TutorialModel
	practice ui.PracticeModel
	width    int
	height   int
}

func initialModel() model {
	return model{
		screen:   screenMenu,
		menu:     ui.NewMenuModel(),
		tutorial: ui.NewTutorialModel(),
		practice: ui.NewPracticeModel(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case ui.SwitchToTutorial:
		m.tutorial = ui.NewTutorialModel()
		m.screen = screenTutorial
		return m, nil
	case ui.SwitchToPractice:
		m.practice = ui.NewPracticeModel()
		m.screen = screenPractice
		return m, nil
	case ui.SwitchToMenu:
		m.screen = screenMenu
		return m, nil
	}

	var cmd tea.Cmd
	switch m.screen {
	case screenMenu:
		m.menu, cmd = m.menu.Update(msg)
	case screenTutorial:
		m.tutorial, cmd = m.tutorial.Update(msg)
	case screenPractice:
		m.practice, cmd = m.practice.Update(msg)
	}
	return m, cmd
}

func (m model) View() string {
	switch m.screen {
	case screenMenu:
		return m.menu.View()
	case screenTutorial:
		return m.tutorial.View()
	case screenPractice:
		return m.practice.View()
	}
	return ""
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
