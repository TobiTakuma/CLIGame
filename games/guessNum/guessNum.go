package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// Model
type model struct {
	targetNum    int
	textInput    textinput.Model
	err          error
	quitting     bool
	stringChoice string
	anounce      string
	attempts     int
	min          int
	max          int
	clear        bool
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "100"
	ti.SetVirtualCursor(false)
	ti.Focus()
	ti.CharLimit = 3
	ti.SetWidth(3)

	return model{
		targetNum: rand.IntN(100),
		min:       0,
		max:       100,
		clear:     false,
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
			if !m.clear {
				m.stringChoice = m.textInput.Value()

				choiceNum, err := strconv.Atoi(m.stringChoice)
				if err != nil {
					return m, nil
				}

				m.attempts++

				if choiceNum == m.targetNum {
					m.anounce = fmt.Sprintf("Well Done!\nTarget Number was %d!\nIt took %v attempts to guess this number.", choiceNum, m.attempts)
					m.clear = true
				} else if choiceNum < m.targetNum {
					if m.min < choiceNum {
						m.min = choiceNum
					}
					m.anounce = fmt.Sprintf("My number is greater than %d", choiceNum)
				} else {
					if m.max > choiceNum {
						m.max = choiceNum
					}
					m.anounce = fmt.Sprintf("My number is less than %d", choiceNum)
				}
			} else {
				return initialModel(), nil
			}
			m.textInput.Reset()
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

// View
func (m model) View() tea.View {
	s := "Start Guessing my number!\n\n"
	s += m.anounce
	s += fmt.Sprintf("\n(%d <= x <= %d)", m.min, m.max)
	if !m.clear {
		s += "\n\n\n\n"
	} else {
		s += "\n\n"
	}
	var c *tea.Cursor
	if !m.textInput.VirtualCursor() {
		c = m.textInput.Cursor()
		c.Y += 7
	}
	s += lipgloss.JoinVertical(lipgloss.Top, m.textInput.View())

	if m.quitting {
		s += "\n"
	}
	s += "\nPress ctrl+c to quit.\n"
	if m.clear {
		s += "Press Enter to Restart\n"
	}

	v := tea.NewView(s)
	v.Cursor = c
	v.AltScreen = true
	return v
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Oof: %v\n", err)
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
