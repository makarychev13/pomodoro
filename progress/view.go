package progress

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

//View описывает оставшееся время помидорки.
type View struct {
	remainMin int
	remainSec int
}

func NewView(remainMin, remainSec int) View {
	return View{
		remainMin: remainMin,
		remainSec: remainSec,
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
		}

	case tickMsg:
		if p.remainSec == 0 {
			p.remainMin--
			p.remainSec = 59
		} else {
			p.remainSec--
		}
		return p, NewTickCommand()
	}

	return p, nil
}

func (p View) View() string {
	return fmt.Sprintf("Осталось %v:%v\n\n%s", p.remainMin, p.remainSec, "Для выхода нажмите q")
}
