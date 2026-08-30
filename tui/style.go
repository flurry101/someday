package tui

import "github.com/charmbracelet/lipgloss"

var (
	titleStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF2A85")).Bold(true)
	dotStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#00F5FF")).Bold(true)
	cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#05FFA1")).Bold(true)
	activeStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#00F5FF")).Bold(true)
	dimStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#6C7A89"))
	pickStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF71CE")).Italic(true)
	statusStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#05FFA1"))
	errStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF3366"))
	helpStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#A06CD5"))
	promptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF2A85")).Bold(true)
)
