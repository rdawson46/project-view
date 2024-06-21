package app

import (
    "github.com/charmbracelet/lipgloss"
)

var (
    wrapper = lipgloss.NewStyle().
        Align(lipgloss.Center, lipgloss.Center).
        Border(lipgloss.RoundedBorder(), true, true, true, true).
        BorderForeground(lipgloss.Color("21"))
)

func (a App) View() string {
    w := wrapper.Width(a.width).Height(a.height).Align(lipgloss.Center)
    return w.Render(a.screen.View())
}
