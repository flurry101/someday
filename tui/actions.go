package tui

import (
	"math/rand"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"someday/db"
)

func (m *model) handleListKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit
	case "j", "down":
		if m.cursor < len(m.items)-1 {
			m.cursor++
		}
		m.picked = -1
	case "k", "up":
		if m.cursor > 0 {
			m.cursor--
		}
		m.picked = -1
	case "a":
		m.mode = modeAdd
		m.input.Focus()
		m.status = ""
		return m, textinput.Blink
	case "d", "enter":
		if len(m.items) > 0 {
			id := m.items[m.cursor].ID
			if err := db.MarkDone(m.db, id); err != nil {
				m.err = err
			} else {
				m.status = "marked done"
			}
			m.reload()
		}
		m.picked = -1
	case "x":
		if len(m.items) > 0 {
			id := m.items[m.cursor].ID
			if err := db.RemoveItem(m.db, id); err != nil {
				m.err = err
			} else {
				m.status = "removed"
			}
			m.reload()
		}
		m.picked = -1
	case "r":
		if len(m.items) > 0 {
			i := rand.Intn(len(m.items))
			m.cursor = i
			m.picked = m.items[i].ID
			m.status = ""
		}
	}
	return m, nil
}
