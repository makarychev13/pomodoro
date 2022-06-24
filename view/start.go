package view

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/makarychev13/pomodoro/command"
)

//Start описывает вид стартового экрана приложения.
type Start struct {
	buttons []period
	cursor  int
}

type period struct {
	minutes int
	text    string
}

func NewStart() Start {
	return Start{
		buttons: []period{
			{minutes: 25, text: "25 мин"},
			{minutes: 15, text: "15 мин"},
			{minutes: 5, text: "5 мин"},
		},
	}
}

func (s Start) Init() tea.Cmd {
	return nil
}

func (s Start) View() string {
	var builder strings.Builder

	builder.WriteString("Запустить таймер?\n\n")

	for i, b := range s.buttons {
		if i == s.cursor {
			builder.WriteString(selectedButtonStyle.Render(b.text))
		} else {
			builder.WriteString(unselectedButtonStyle.Render(b.text))
		}
	}

	return mainWindowStyle.Render(builder.String())
}

func (s Start) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "й":
			return s, tea.Quit
		case tea.KeyLeft.String():
			s.cursor--
			if s.cursor < 0 {
				s.cursor = len(s.buttons) - 1
			}
		case tea.KeyRight.String():
			s.cursor++
			if s.cursor == len(s.buttons) {
				s.cursor = 0
			}
		case tea.KeyEnter.String(), tea.KeySpace.String(), tea.KeyRight.String():
			return NewProgress(s.buttons[s.cursor].minutes, 0), command.NewTick()
		}
	}

	return s, nil
}
