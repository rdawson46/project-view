package app

import (
    tea "github.com/charmbracelet/bubbletea"
)

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    return a, nil
}
