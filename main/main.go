package main

import (
	"fmt"
	"os"
	"sqlite-tui-go/database"
	"sqlite-tui-go/views"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	d := database.InitDatabase("test.db")

	p := tea.NewProgram(views.NewModel(d), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("something went wrong!, %v", err)
		os.Exit(1)
	}
}
