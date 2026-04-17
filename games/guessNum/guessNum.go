package main

import (
	"fmt"
	"os"
	"strconv"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// Model
type model struct {
	textInput textinput.Model
	err       error
	quitting  bool
	choice    int
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "hello"
	ti.SetVirtualCursor(false)
	ti.Focus()
	ti.CharLimit = 156
	ti.SetWidth(30)

	return model{
		textInput: ti,
	}
}

// Init
func (m model) Init() tea.Cmd {
	return textinput.Blink
}

// Update
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			val := m.textInput.Value()

			i, err := strconv.Atoi(val)
			if err != nil {
				return m, nil
			}

			m.choice = i
			return m, nil
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

// View
func (m model) View() tea.View {
	s := "Start Guessing my number!\n\n You entered: ", m.choice

	var c *tea.Cursor
	if !m.textInput.VirtualCursor() {
		c = m.textInput.Cursor()
		c.Y += 2
	}
	s += lipgloss.JoinVertical(lipgloss.Top, m.textInput.View())

	if m.quitting {
		s += "\n"
	}
	s += "\nPress ctrl+c to quit.\n"

	v := tea.NewView(s)
	v.Cursor = c
	v.AltScreen = true
	return v
}

func main() {
	p := tea.NewProgram(initialModel())
	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	if m, ok := m.(model); ok && m.choice != 0 {
		fmt.Printf("\n---\nYou typed ", m.choice)
	}

	// targetNum := (rand.IntN(100))
	// fmt.Println("Start guessing my number!")
	// var attempts int = 0
	//
	//	for {
	//		attempts++
	//		fmt.Print("Type a number: ")
	//		var i int
	//
	//		fmt.Scan((&i))
	//		if i == targetNum {
	//			fmt.Printf("Well Done! It took %v attempts to guess this number.", attempts)
	//			break
	//		} else if i < targetNum {
	//			fmt.Println("My number is greater than", i)
	//		} else {
	//			fmt.Println("My number is less than", i)
	//		}
	//	}
}
