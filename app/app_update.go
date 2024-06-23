package app

import (
    tea "github.com/charmbracelet/bubbletea"
)

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        a.height = msg.Height - 2
        a.width = msg.Width - 2

    case tea.KeyMsg:
        str := msg.String()

        switch str {
        case "q", "ctrl+c":
            return a, tea.Quit
        case "?":
            a = a.changeScreen(Help)
        case "f":
            a = a.changeScreen(Finder)
        case "e":
            a = a.changeScreen(File)
        case "p":
            a = a.changeScreen(Explorer)
        }
    }   

    return a, nil
}
