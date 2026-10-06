package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

// model holds all application state. In bubbletea's Elm architecture
// state lives in a single struct and is only ever changed in Update.
type model struct {
	count int
}

// Init runs once at startup. Commands (tea.Cmd) perform asynchronous work
// such as timers or I/O; this example needs none, so it returns nil.
func (m model) Init() tea.Cmd {
	return nil
}

// Update receives messages (key presses, completed commands, ...) and
// returns a new model, plus optionally a command to run. The value
// receiver means state changes are made by returning a modified copy.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "+", "up":
			m.count++
		case "-", "down":
			m.count--
		}
	}
	return m, nil
}

// View renders the current state. bubbletea repaints the terminal with
// whatever View returns after each Update. In bubbletea v2 the view is a
// tea.View (built with tea.NewView), not a plain string.
func (m model) View() tea.View {
	return tea.NewView(fmt.Sprintf(`Counter: %d

+ / up    increment
- / down  decrement
q / ctrl+c  quit
`, m.count))
}

func main() {
	final, err := tea.NewProgram(model{}).Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if m, ok := final.(model); ok {
		fmt.Printf("Final counter value: %d\n", m.count)
	}
}
