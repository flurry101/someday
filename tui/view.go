package tui

import (
	"fmt"
	"strings"
)

func (m *model) View() string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("someday") + dotStyle.Render("."))
	b.WriteString("\n\n")

	if len(m.items) == 0 {
		b.WriteString(dimStyle.Render("nothing here yet. press 'a' to add something.") + "\n")
	}

	for i, it := range m.items {
		cursor := "  "
		line := it.Text
		if i == m.cursor {
			cursor = cursorStyle.Render("› ")
			line = activeStyle.Render(it.Text)
		} else {
			line = dimStyle.Render(it.Text)
		}
		mark := ""
		if m.picked == it.ID {
			mark = pickStyle.Render(" ← someday")
		}
		fmt.Fprintf(&b, "%s%s%s\n", cursor, line, mark)
	}

	b.WriteString("\n")
	if m.mode == modeAdd {
		b.WriteString(promptStyle.Render("add: ") + m.input.View() + "\n")
	} else if m.status != "" {
		b.WriteString(statusStyle.Render(m.status) + "\n")
	} else if m.err != nil {
		b.WriteString(errStyle.Render("error: "+m.err.Error()) + "\n")
	} else {
		b.WriteString("\n")
	}

	b.WriteString(helpStyle.Render("a add   d done   x remove   r random   j/k move   q quit"))
	return b.String()
}
