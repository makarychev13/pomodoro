package start

import "github.com/charmbracelet/lipgloss"

var (
	mainWindowStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#171717")).
			Foreground(lipgloss.Color("#bebebe")).
			Width(44).
			PaddingTop(1).
			PaddingBottom(1).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("63")).
			BorderBackground(lipgloss.Color("#171717")).
			Align(lipgloss.Center)

	selectedItemStyle = lipgloss.NewStyle().
				Background(lipgloss.Color("63")).
				Foreground(lipgloss.Color("#bebebe")).
				MarginLeft(1).
				MarginRight(1).
				PaddingLeft(1).
				PaddingRight(1).
				MarginBackground(lipgloss.Color("#171717")).
				Align(lipgloss.Center)

	unselectedItemStyle = lipgloss.NewStyle().
				Background(lipgloss.Color("#878B7D")).
				Foreground(lipgloss.Color("bebebe")).
				MarginLeft(1).
				MarginRight(1).
				PaddingLeft(1).
				PaddingRight(1).
				MarginBackground(lipgloss.Color("#171717")).
				Align(lipgloss.Center)
)
