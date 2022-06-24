package command

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	timeout = time.Second
)

//TickMsg описывает сообщения команды.
type TickMsg time.Time

//NewTick создаёт команду, которая по таймеру отправляет сообщение об уменьшении оставшегося времени помидорки.
func NewTick() tea.Cmd {
	return func() tea.Msg {
		return tick(timeout)
	}
}

func tick(duration time.Duration) tea.Msg {
	n := time.Now()
	d := n.Truncate(duration).Add(duration).Sub(n)
	t := time.NewTimer(d)

	return TickMsg(<-t.C)
}
