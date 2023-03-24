package main

import (
	"fmt"
	"os"
	"sqlite-tui-go/views"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(views.InitialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("something went wrong!, %v", err)
		os.Exit(1)
	}
}
