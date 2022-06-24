package start

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/makarychev13/pomodoro/progress"
)

//View описывает вид стартового экрана приложения.
type View struct {
	ranges []period
	cursor int
}

type period struct {
	minutes int
	text    string
}

func NewView() View {
	return View{
		ranges: []period{
			{minutes: 25, text: "25 мин"},
			{minutes: 15, text: "15 мин"},
			{minutes: 5, text: "5 мин"},
		},
	}
}

func (s View) Init() tea.Cmd {
	return nil
}

func (s View) View() string {
	var builder strings.Builder

	builder.WriteString("Запустить таймер?\n\n")

	for i, r := range s.ranges {
		if i == s.cursor {
			builder.WriteString(selectedItemStyle.Render(r.text))
		} else {
			builder.WriteString(unselectedItemStyle.Render(r.text))
		}
	}

	return mainWindowStyle.Render(builder.String())
}

func (s View) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "й":
			return s, tea.Quit
		case tea.KeyLeft.String():
			s.cursor--
			if s.cursor < 0 {
				s.cursor = len(s.ranges) - 1
			}
		case tea.KeyRight.String():
			s.cursor++
			if s.cursor == len(s.ranges) {
				s.cursor = 0
			}
		case tea.KeyEnter.String(), tea.KeySpace.String(), tea.KeyRight.String():
			return progress.NewView(s.ranges[s.cursor].minutes, 0), progress.NewTickCommand()
		}
	}

	return s, nil
}
