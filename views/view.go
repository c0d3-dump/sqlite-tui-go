package views

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	return query_view()
}

func table_view() string {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			top_section.Render("test_table"),
			button_style.Render("SUBMIT"),
		),
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			left_section.Render("ids"),
			right_section.Render("data"),
		),
	)
}

func create_table_view() string {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			top_section.Render("test_table"),
			button_style.Render("SUBMIT"),
		),
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			left_section.Render("column name"),
			right_section.Render("column data"),
		),
	)
}

func create_column_view() string {
	return base_section.Render(
		lipgloss.JoinVertical(
			lipgloss.Bottom,
			create_column_section.Render("create column"),
			button_style.Render("SUBMIT"),
		),
	)
}

func select_table_view() string {
	return base_section.Render(
		create_column_section.Render("select table"),
	)
}

func query_view() string {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			query_section.Render(""),
			button_style.Render("SUBMIT"),
		),
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			left_section.Render("index"),
			right_section.Render("data"),
		),
	)
}
