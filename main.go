package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

const url = "https://charm.sh/"

type responseStruct struct {
	City       string   `json:"city"`
	RegionCode string   `json:"region_code"`
	OS         string   `json:"os"`
	Isp        string   `json:"isp"`
	Ports      []int    `json:"ports"`
	Hostnames  []string `json:"hostnames"`
	Data       []struct {
		Timestamp string `json:"timestamp"`
	} `json:"data"`
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type (
	errMsg error
)

type model struct {
	textInput     textinput.Model
	err           error
	showPrompt    bool // Flag to control the display of the initial prompt
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Type a Pokémon name"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		textInput: ti,
		err:       nil,
		showPrompt: true, // Initially set to true to show the prompt
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			// Handle different inputs
			switch m.textInput.Value() {
			case "Pikachu":
				handlePikachu()
			case "Raichu":
				handleRaichu()
			}
			// Reset text input for next input and hide the initial prompt
			m.textInput.SetValue("")
			m.textInput.Focus()
			m.showPrompt = false // Do not show the initial prompt again
			return m, nil
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// Handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	var view string
	if m.showPrompt {
		view = fmt.Sprintf(
			"What’s your favorite Pokémon?\n\n%s\n\n%s",
			m.textInput.View(),
			"(Enter to confirm, ESC to quit)",
		)
	} else {
		view = fmt.Sprintf(
			"%s\n\n%s",
			m.textInput.View(),
			"(Enter to confirm, ESC to quit)",
		)
	}
	return view + "\n"
}

// Define your custom functions for handling specific inputs
func handlePikachu() {
	fmt.Println("Pikachu function is called")
	// Add your logic here
}

func handleRaichu() {
	fmt.Println("Raichu function is called")
	// Add your logic here
}
