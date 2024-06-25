package app

import (
    tea "github.com/charmbracelet/bubbletea"
)

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch message := msg.(type) {
    case tea.WindowSizeMsg:
        a.height = message.Height - 2
        a.width = message.Width - 2

        var cmd tea.Cmd
        a.screen, cmd = a.screen.Update(msg)
        return a, cmd

    case tea.KeyMsg:
        str := message.String()

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
        default:
            var cmd tea.Cmd
            a.screen, cmd = a.screen.Update(msg)
            return a, cmd
        }
    default:
        var cmd tea.Cmd
        a.screen, cmd = a.screen.Update(msg)
        return a, cmd
    }   

    return a, nil
}
