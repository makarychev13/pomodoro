package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/makarychev13/pomodoro/view"
)

func main() {
	p := tea.NewProgram(view.NewStart())

	if err := p.Start(); err != nil {
		fmt.Printf("Не удалось запустить программу: %v", err)
		os.Exit(1)
	}
}
