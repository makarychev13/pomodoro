package view

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	not_selected = -1
)

//Start описывает вид стартового экрана приложения.
type Start struct {
	ranges   []string
	cursor   int
	selected int
}

func NewStart() Start {
	return Start{
		ranges:   []string{"25 мин", "15 мин", "5 мин"},
		selected: not_selected,
	}
}

func (s Start) Init() tea.Cmd {
	return nil
}

func (s Start) View() string {
	var builder strings.Builder

	builder.WriteString("Запустить таймер?\n\n")

	for i, r := range s.ranges {
		cursor := " "
		if i == s.cursor {
			cursor = ">"
		}

		selected := "[ ]"
		if i == s.selected {
			selected = "[x]"
		}

		builder.WriteString(fmt.Sprintf("%s %s %s\n", cursor, selected, r))
	}

	builder.WriteString("\nДля выхода нажмите q")

	return builder.String()
}

func (s Start) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "й":
			return s, tea.Quit
		case tea.KeyUp.String():
			s.cursor--
			if s.cursor < 0 {
				s.cursor = len(s.ranges) - 1
			}
		case tea.KeyDown.String():
			s.cursor++
			if s.cursor == len(s.ranges) {
				s.cursor = 0
			}
		case tea.KeyEnter.String(), tea.KeySpace.String(), tea.KeyRight.String():
			return NewProgress(s.ranges[s.cursor]), nil
		case tea.KeyLeft.String():
			s.selected = not_selected
		}
	}

	return s, nil
}
