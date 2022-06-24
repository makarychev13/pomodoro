package view

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/makarychev13/pomodoro/command"
)

const (
	pause   = "Пауза"
	exit    = "Выйти"
	run     = "Продолжить"
	restart = "Заново"
)

//Progress описывает оставшееся время помидорки.
type Progress struct {
	remainMin int
	remainSec int
	startMin  int
	startSec  int
	cursor    int
	buttons   []string
	active    bool
}

func NewProgress(startMin, startSec int) Progress {
	return Progress{
		startMin:  startMin,
		startSec:  startSec,
		remainMin: startMin,
		remainSec: startSec,
		buttons:   []string{pause, restart, exit},
		cursor:    0,
		active:    true,
	}
}

func (p Progress) Init() tea.Cmd {
	return nil
}

func (p Progress) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

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
					p.active = true
					cmd = command.NewTick()
				}
			}

			if p.cursor == 1 {
				p.remainMin = p.startMin
				p.remainSec = p.startSec
			}

			if p.cursor == 2 {
				return NewStart(), nil
			}
		}

	case command.TickMsg:
		if p.remainMin == 0 && p.remainSec == 0 {
			return NewPause(1, 0, p.startMin), command.NewTick()
		}

		if p.active {
			if p.remainSec == 0 {
				p.remainMin--
				p.remainSec = 59
			} else {
				p.remainSec--
			}

			cmd = command.NewTick()
		}
	}

	return p, cmd
}

func (p Progress) View() string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("Осталось %d:%02d\n\n", p.remainMin, p.remainSec))

	for i, b := range p.buttons {
		if i == p.cursor {
			builder.WriteString(selectedButtonStyle.Render(b))
		} else {
			builder.WriteString(unselectedButtonStyle.Render(b))
		}
	}

	return mainWindowStyle.Render(builder.String())
}
