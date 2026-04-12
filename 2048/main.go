package main

import (
	// go mod init 2048
	// go get github.com/charmbracelet/bubbletea
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "w":

		}
	}
	return m, nil
}

func (m model) View() string {
	return "Hello"
}

type model struct{}

func main() {
	p := tea.NewProgram(model{})
	p.Run()
}
