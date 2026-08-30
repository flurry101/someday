package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"someday/db"
)

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.mode == modeAdd {
			return m.handleAddKey(msg)
		}
		return m.handleListKey(msg)
	}
	return m, nil
}

func (m *model) handleAddKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyEnter:
		text := strings.TrimSpace(m.input.Value())
		if text != "" {
			if _, err := db.AddItem(m.db, text); err != nil {
				m.err = err
			} else {
				m.status = "added"
			}
		}
		m.input.SetValue("")
		m.mode = modeList
		m.reload()
		return m, nil
	case tea.KeyEsc:
		m.input.SetValue("")
		m.mode = modeList
		return m, nil
	}
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}
