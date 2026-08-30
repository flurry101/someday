package tui

import (
	"database/sql"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"someday/db"
)

type mode int

const (
	modeList mode = iota
	modeAdd
)

type model struct {
	db     *sql.DB
	items  []db.Item
	cursor int
	mode   mode
	input  textinput.Model
	status string
	picked int
	err    error
}

func Run(database *sql.DB) error {
	ti := textinput.New()
	ti.Placeholder = "something you want to do someday"
	ti.CharLimit = 200

	m := model{db: database, input: ti, picked: -1}
	if err := m.reload(); err != nil {
		return err
	}
	p := tea.NewProgram(&m, tea.WithAltScreen())
	_, err := p.Run()
	return err
}

func (m *model) reload() error {
	items, err := db.ListItems(m.db, false)
	if err != nil {
		return err
	}
	m.items = items
	if m.cursor >= len(m.items) {
		m.cursor = len(m.items) - 1
	}
	if m.cursor < 0 {
		m.cursor = 0
	}
	return nil
}

func (m *model) Init() tea.Cmd {
	return nil
}
