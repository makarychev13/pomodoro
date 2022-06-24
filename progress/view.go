package progress

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/makarychev13/pomodoro/style"
)

const (
	pause   = "Пауза"
	exit    = "Выйти"
	run     = "Продолжить"
	restart = "Заново"
)

//View описывает оставшееся время помидорки.
type View struct {
	remainMin int
	remainSec int
	cursor    int
	buttons   []string
	active    bool
}

func NewView(remainMin, remainSec int) View {
	return View{
		remainMin: remainMin,
		remainSec: remainSec,
		buttons:   []string{pause, restart, exit},
		cursor:    0,
		active:    true,
	}
}

func (p View) Init() tea.Cmd {
	return nil
}

func (p View) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "й":
			return p, tea.Quit
		case tea.KeyRight.String():
			p.cursor++
			if p.cursor == len(p.buttons) {
				p.cursor = 0
			}
		case tea.KeyLeft.String():
			p.cursor--
			if p.cursor < 0 {
				p.cursor = len(p.buttons) - 1
			}
		case tea.KeyEnter.String(), tea.KeySpace.String():
			if p.cursor == 0 {
				if p.buttons[0] == pause {
					p.buttons[0] = run
					p.active = false
				} else if p.buttons[0] == run {
					p.buttons[0] = pause
				}
			}
		}

	case tickMsg:
		var cmd tea.Cmd

		if p.active {
			if p.remainSec == 0 {
				p.remainMin--
				p.remainSec = 59
			} else {
				p.remainSec--
			}

			cmd = NewTickCommand()
		}

		return p, cmd
	}

	return p, nil
}

func (p View) View() string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("Осталось %v:%v\n\n", p.remainMin, p.remainSec))

	for i, b := range p.buttons {
		if i == p.cursor {
			builder.WriteString(style.SelectedButton.Render(b))
		} else {
			builder.WriteString(style.UnselectedButton.Render(b))
		}
	}

	return style.MainWindow.Render(builder.String())
}
