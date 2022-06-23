package progress

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	timeout = time.Second
)

type tickMsg time.Time

func NewTickCommand() tea.Cmd {
	return func() tea.Msg {
		return tick(timeout)
	}
}

func tick(duration time.Duration) tea.Msg {
	n := time.Now()
	d := n.Truncate(duration).Add(duration).Sub(n)
	t := time.NewTimer(d)

	return tickMsg(<-t.C)
}
