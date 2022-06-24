package view

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/makarychev13/pomodoro/command"
)

//Pause описывает экран маленькой помидорки отдыха.
type Pause struct {
	remainMin    int
	remainSec    int
	progressInit int
	cursor       int
	buttons      []prolongation
	active       bool
}

type prolongation struct {
	text   string
	period int
}

func NewPause(startMin, startSec, progressInit int) Pause {
	return Pause{
		remainMin:    startMin,
		remainSec:    startSec,
		progressInit: progressInit,
		buttons: []prolongation{
			{
				text:   "+2 мин",
				period: 2,
			},
			{
				text:   "+5 мин",
				period: 5,
			},
			{
				text:   "Пропустить",
				period: 0,
			},
		},
		active: true,
	}
}

func (p Pause) Init() tea.Cmd {
	return nil
}

func (p Pause) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if p.cursor != 2 {
				p.remainMin += p.buttons[p.cursor].period
			} else {
				return NewProgress(p.progressInit, 0), nil
			}
		}

	case command.TickMsg:
		if p.remainMin == 0 && p.remainSec == 0 {
			return NewProgress(p.progressInit, 0), command.NewTick()
		}

		if p.remainSec == 0 {
			p.remainMin--
			p.remainSec = 59
		} else {
			p.remainSec--
		}

		cmd = command.NewTick()
	}
	return p, cmd
}

func (p Pause) View() string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("Перерыв %d:%02d\n\n", p.remainMin, p.remainSec))

	for i, b := range p.buttons {
		if i == p.cursor {
			builder.WriteString(selectedButtonStyle.Render(b.text))
		} else {
			builder.WriteString(unselectedButtonStyle.Render(b.text))
		}
	}

	return mainWindowStyle.Render(builder.String())
}
