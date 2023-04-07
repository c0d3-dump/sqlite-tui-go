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
	var cmd tea.Cmd
	var boolVal bool

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "ctrl+t":
			m.cursor[0] = 1
			m.cursor[1] = 0
			m.cursor[2] = 0
			m.textInput.SetValue("@")
		case "ctrl+n":
			if m.cursor[0] == 1 {
				m.cursor[1] = 1
				m.cursor[2] = 0
				m.textInput.SetValue("@")
			}
		case "ctrl+q":
			m.cursor[0] = 2
			m.cursor[1] = 0
			m.cursor[2] = 0
			m.textInput.SetValue("@")
		case "ctrl+o":
			if m.cursor[0] == 2 {
				m.cursor[1] = 1
				m.cursor[2] = 0
				m.textInput.SetValue("@")
			}
		case "ctrl+a":
			if m.cursor[0] == 2 {
				m.cursor[1] = 2
				m.cursor[2] = 0
				m.textInput.SetValue("@")
			}
		case "down":
			if m.cursor[0] == 1 && m.cursor[1] == 1 && m.cursor[2] == 0 {
				m.createColumn.cursor = (m.createColumn.cursor + 1) % 4
				m.textInput.SetValue("@")
			} else if m.cursor[0] == 1 && m.cursor[1] == 0 && m.cursor[2] == 1 {
				m.createTable.cursor = (m.createTable.cursor + 1) % len(m.createTable.columns)
			} else if m.cursor[0] == 2 && m.cursor[1] == 0 && m.cursor[2] == 1 {
				m.tables[0].cursor = (m.tables[0].cursor + 1) % len(m.tables[0].data)
			} else if m.cursor[0] == 2 && m.cursor[1] == 1 {
				m.currentTable = (m.currentTable + 1) % len(m.tables)
			} else if m.cursor[0] == 2 && m.cursor[1] == 2 && m.cursor[2] == 0 {
				m.tables[m.currentTable].addRow.cursor = (m.tables[m.currentTable].addRow.cursor + 1) % (len(m.tables[m.currentTable].columns) - 1)
				m.textInput.SetValue("@")
			}
		case "up":
			if m.cursor[0] == 1 && m.cursor[1] == 1 && m.cursor[2] == 0 {
				m.createColumn.cursor = (m.createColumn.cursor - 1) % 4
				if m.createColumn.cursor < 0 {
					m.createColumn.cursor = 3
				}
				m.textInput.SetValue("@")
			} else if m.cursor[0] == 1 && m.cursor[1] == 0 && m.cursor[2] == 1 {
				m.createTable.cursor = (m.createTable.cursor - 1) % len(m.createTable.columns)
				if m.createTable.cursor < 0 {
					m.createTable.cursor = len(m.createTable.columns) - 1
				}
			} else if m.cursor[0] == 2 && m.cursor[1] == 0 && m.cursor[2] == 1 {
				m.tables[0].cursor = (m.tables[0].cursor - 1) % len(m.tables[0].data)
				if m.tables[0].cursor < 0 {
					m.tables[0].cursor = len(m.tables[0].data) - 1
				}
			} else if m.cursor[0] == 2 && m.cursor[1] == 1 {
				m.currentTable = (m.currentTable - 1) % len(m.tables)
				if m.currentTable < 0 {
					m.currentTable = len(m.tables) - 1
				}
			} else if m.cursor[0] == 2 && m.cursor[1] == 2 && m.cursor[2] == 0 {
				m.tables[m.currentTable].addRow.cursor = (m.tables[m.currentTable].addRow.cursor - 1) % (len(m.tables[m.currentTable].columns) - 1)
				if m.tables[m.currentTable].addRow.cursor < 0 {
					m.tables[m.currentTable].addRow.cursor = len(m.tables[m.currentTable].columns) - 2
				}
				m.textInput.SetValue("@")
			}
		case "right":
			boolVal = true
		case "left":
			boolVal = false
		case "tab":
			switch m.cursor[0] {
			case 1:
				switch m.cursor[1] {
				case 0:
					m.cursor[2] = (m.cursor[2] + 1) % 4
				case 1:
					m.cursor[2] = (m.cursor[2] + 1) % 2
				}
			case 2:
				switch m.cursor[1] {
				case 0:
					m.cursor[2] = (m.cursor[2] + 1) % 4
				case 2:
					m.cursor[2] = (m.cursor[2] + 1) % 2
				}
			}
			m.textInput.SetValue("@")
		case "enter":
			switch m.cursor[0] {
			case 1:
				switch m.cursor[1] {
				case 0:
					switch m.cursor[2] {
					case 3:
						m.CreateTable(m.createTable)
						m.createTable = CreateTable{}
						m.UpdateTable()
					}
				case 1:
					switch m.cursor[2] {
					case 1:
						m.createTable.columns = append(m.createTable.columns, m.createColumn)
						m.createColumn = CreateColumn{}
					}
				}
			case 2:
				switch m.cursor[1] {
				case 1:
					m.cursor[1] = 0
				case 2:
					switch m.cursor[2] {
					case 1:
						m.AddRow()
						m.tables[m.currentTable].addRow = AddRow{
							cursor: 0,
							data:   make([]string, len(m.tables[m.currentTable].columns)-1),
						}
						m.UpdateTable()
					}
				}

			}
			m.textInput.SetValue("@")
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
	}

	m.textInput, cmd = m.textInput.Update(msg)

	switch m.cursor[0] {
	case 1:
		switch m.cursor[1] {
		case 0:
			m.textInput.Width = query_section.GetWidth() - width_offset*2
			if m.textInput.Value() != "@" {
				m.createTable.name = m.textInput.Value()
			}
			m.textInput.SetValue(m.createTable.name)
			m.textInput.CursorEnd()
		case 1:
			m.textInput.Width = create_column_section.GetWidth() - width_offset*3
			switch m.createColumn.cursor {
			case 0:
				if m.textInput.Value() != "@" {
					m.createColumn.name = m.textInput.Value()
				}
				m.textInput.SetValue(m.createColumn.name)
			case 1:
				if m.textInput.Value() != "@" {
					m.createColumn.dtype = m.textInput.Value()
				}
				m.textInput.SetValue(m.createColumn.dtype)
			case 2:
				if m.textInput.Value() != "@" {
					m.createColumn.notnull = boolVal
				}
				var val string
				if m.createColumn.notnull {
					val = "true"
				} else {
					val = "false"
				}
				m.textInput.SetValue(val)
			case 3:
				if m.textInput.Value() != "@" {
					m.createColumn.dval = m.textInput.Value()
				}
				m.textInput.SetValue(m.createColumn.dval)
			}
			m.textInput.CursorEnd()
		}
	case 2:
		switch m.cursor[1] {
		case 0:
			if m.currentTable < 0 {
				m.cursor[1] = 1
			}
			m.textInput.Width = query_section.GetWidth() - width_offset*2
			if m.textInput.Value() != "@" {
				m.queryText = m.textInput.Value()
			}
			m.textInput.SetValue(m.queryText)
			m.textInput.CursorEnd()
		case 2:
			m.textInput.Width = create_column_section.GetWidth() - width_offset*3
			addRow := m.tables[m.currentTable].addRow
			if m.textInput.Value() != "@" {
				m.tables[m.currentTable].addRow.data[addRow.cursor] = m.textInput.Value()
			}
			m.textInput.SetValue(addRow.data[addRow.cursor])
			m.textInput.CursorEnd()
		}
	}

	return m, cmd
}
