package main

import (
	"fmt"
	"os"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// Model
type model struct {
	cursor   int
	choices  []string
	selected map[int]struct{}

	textInput textinput.Model
	err       error
	quitting  bool
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "50"
	ti.SetVirtualCursor(false)
	ti.Focus()
	ti.CharLimit = 156

	return model{
		choices:    []string{"hello", "takuma", "toiyama", "aiueo", ";aldkjf"},
		selected:   make(map[int]struct{}),
		texitInput: ti,
	}
}

// Init
func (m model) Init() tea.Cmd {
	return textinput.Blink
}

// Update
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", "space":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

// View
func (m model) View() tea.View {
	s := "Start Guessing my number!\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	var c *tea.Cursor
	if !m.textInputVirtualCursor() {
		c = m.textInput.Cursor
		c.Y += lipgloss.Height(m.headerView())
	}

	s += lipgloss.JoinVertical(lipgloss.Top, m.headerView(), m.textInput.View(), m.footerVieq())
	if m.quitting {
		s += "\n"
	}

	s += "\nPress q to quit.\n"
	s += "Hello World!\n"

	v := tea.NewView(s)
	v.WindowTitle = "Grocery List"

	return v
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
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
