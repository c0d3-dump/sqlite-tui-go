package views

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

var (
	columnStyle = lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62"))
)

func (m Model) View() string {
	termenv.Size()

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		lipgloss.JoinVertical(
			lipgloss.Left,
			columnStyle.Render("Hello, world"),
			columnStyle.Render("Hello, world"),
		),
		columnStyle.Render("Hello, world"),
	)
}
