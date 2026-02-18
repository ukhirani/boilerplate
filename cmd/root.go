package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
	"github.com/ukhirani/boilerplate/constants"
	"github.com/ukhirani/boilerplate/services"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor   int
	choices  []string
	selected map[int]struct{}
}

func initialModel() model {
	allTemplates, _ := services.GetAllTemplates()
	return model{
		choices: allTemplates,

		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("Grocery List")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
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
		case "enter", " ":
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

func (m model) View() string {
	s := "What templates do you wanna clone ? \n\n"

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

	s += "\nPress q to quit.\n"

	return s
}

var rootCmd = &cobra.Command{
	Use:   "bp",
	Short: "A CLI tool to ironically skip writing boilerplate code",
	Long: `Manage reusable file and directory templates.
Save common code structures as templates and generate them instantly.`,
	Run: RootRunner,
}

func RootRunner(cmd *cobra.Command, args []string) {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func Execute() {
	services.InitViper()

	// overriding fang's versioning
	rootCmd.SetVersionTemplate(constants.VERSION)
	rootCmd.Version = constants.VERSION

	if err := fang.Execute(context.Background(), rootCmd); err != nil {
		os.Exit(1)
	}
}
