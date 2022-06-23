package view

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

//Progress описывает оставшееся время помидорки
type Progress struct {
	remain string
}

func NewProgress(remain string) Progress {
	return Progress{remain: remain}
}

func (p Progress) Init() tea.Cmd {
	return nil
}

func (p Progress) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "й":
			return p, tea.Quit
		}
	}

	return p, nil
}

func (p Progress) View() string {
	return fmt.Sprintf("Осталось %s\n\n%s", p.remain, "Для выхода нажмите q")
}
