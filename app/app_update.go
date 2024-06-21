package app

import (
    tea "github.com/charmbracelet/bubbletea"
    "strconv"
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
        default:
            x, err := strconv.Atoi(str)

            if err != nil {
                break
            }

            a = a.changeScreen(x)
        }
    }   

    return a, nil
}
