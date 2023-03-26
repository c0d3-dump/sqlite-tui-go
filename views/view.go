package views

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	switch m.cursor[0] {
	case 0:
		return base_view()
	case 1:
		switch m.cursor[1] {
		case 0:
			return create_table_view(m)
		case 1:
			return create_column_view(m)
		}
	case 2:
		switch m.cursor[1] {
		case 0:
			return table_view()
		case 1:
			return select_table_view()
		case 2:
			return add_row_view()
		}
	case 3:
		return query_view(m)
	}
	return base_view()
}

func base_view() string {
	return base_section.Render(
		lipgloss.NewStyle().
			AlignHorizontal(lipgloss.Left).
			Render("There are no tables in db,\nPlease create one by pressing c\n - c : create table\n - v : view table data\n - / : query tables\n"),
	)
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

func select_table_view() string {
	return base_section.Render(
		create_column_section.Render("select table"),
	)
}

func add_row_view() string {
	return base_section.Render(
		create_column_section.Render("add row"),
	)
}

func create_table_view(m Model) string {
	var t string

	if m.cursor[2] == 0 {
		t = m.textInput.View()
	} else {
		t = m.createTable.name
	}

	highlight(m.cursor[2], []lipgloss.Style{
		top_section,
		left_section,
		right_section,
		button_style,
	})

	var colName []string
	var colData []string

	for i, col := range m.createTable.columns {
		if m.createTable.cursor == i {
			colName = append(colName, ">"+col.name)
			colData = append(colData, "data type: "+col.dtype)
			colData = append(colData, "not null: "+boolToString(col.notnull))
			colData = append(colData, "default value: "+col.dval)
			colData = append(colData, "primary key: "+boolToString(col.pk))
		} else {
			colName = append(colName, col.name)
		}
	}

	return lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			top_section.Render(t),
			button_style.Render("SUBMIT"),
		),
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			left_section.Render(
				lipgloss.JoinVertical(
					lipgloss.Top,
					colName...,
				),
			),
			right_section.Render(
				lipgloss.JoinVertical(
					lipgloss.Top,
					colData...,
				),
			),
		),
	)
}

func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func create_column_view(m Model) string {
	t := createColumnListSelect(m, []string{
		"name",
		"dtype",
		"notnull (true/false)",
		"dval",
		"pk (true/false)",
	})

	highlight(m.cursor[2], []lipgloss.Style{
		create_column_section,
		button_style,
	})

	return base_section.Render(
		lipgloss.JoinVertical(
			lipgloss.Bottom,
			create_column_section.Render(t),
			button_style.Render("SUBMIT"),
		),
	)
}

func query_view(m Model) string {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			query_section.Render(
				m.textInput.View(),
			),
			button_style.Render("SUBMIT"),
		),
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			left_section.Render("index"),
			right_section.Render("data"),
		),
	)
}

func highlight(focus int, styles []lipgloss.Style) {
	for i, style := range styles {
		if focus == i {
			style.BorderForeground(lipgloss.Color("#00ff00"))
		} else {
			style.BorderForeground(lipgloss.Color("62"))
		}
	}
}

func createColumnListSelect(m Model, keys []string) string {
	t := [5]string{}

	for i := range keys {
		if m.createColumn.cursor == i {
			t[i] = m.textInput.View()
		} else {
			switch i {
			case 0:
				t[0] = m.createColumn.name
			case 1:
				t[1] = m.createColumn.dtype
			case 2:
				t[2] = boolToString(m.createColumn.notnull)
			case 3:
				t[3] = m.createColumn.dval
			case 4:
				t[4] = boolToString(m.createColumn.pk)
			}
		}
	}

	var temp []string

	for i, k := range keys {
		temp = append(temp, k)
		if t[i] != "" {
			temp = append(temp, text_style.Render(t[i]))
		}
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		temp...,
	)
}
