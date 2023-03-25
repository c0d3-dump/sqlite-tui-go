package views

import (
	tea "github.com/charmbracelet/bubbletea"
)

const (
	height_offset  = 5
	width_offset   = 2
	section_offset = 4
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		top_section.Width(msg.Width - width_offset*2 - button_style.GetWidth())
		query_section.Width(msg.Width - width_offset*2 - button_style.GetWidth())
		left_section.Height(msg.Height - height_offset)
		left_section.Width(msg.Width / section_offset)
		right_section.Height(msg.Height - height_offset)
		right_section.Width(msg.Width - left_section.GetWidth() - width_offset*2)
		base_section.Height(msg.Height)
		base_section.Width(msg.Width)
		create_column_section.Width(msg.Width / 2)
		create_column_section.Height(msg.Height / 2)
		return m, nil
	}
	return m, nil
}
