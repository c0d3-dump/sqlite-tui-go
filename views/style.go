package views

import "github.com/charmbracelet/lipgloss"

var (
	top_section = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62")).
			AlignHorizontal(lipgloss.Center)

	left_section = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62")).
			AlignHorizontal(lipgloss.Center)

	right_section = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62")).
			Padding(0, 1)

	base_section = lipgloss.NewStyle().
			Align(lipgloss.Center, lipgloss.Center)

	create_column_section = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("62")).
				Padding(0, 1)

	button_style = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62")).
			Padding(0, 2).
			Foreground(lipgloss.Color("#00ff00")).
			Width(10).
			Bold(false)

	query_section = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62"))
)
