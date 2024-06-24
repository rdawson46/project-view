package explorer

import tea "github.com/charmbracelet/bubbletea"

func (e Explorer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch message := msg.(type) {
    case tea.KeyMsg:
        switch e.exist {
        case false:
            return e.newPromptUpdate(message.String())
        default:
            return e, nil
        }
    }
    return e, nil
}

func (e Explorer) newPromptUpdate(choice string) (tea.Model, tea.Cmd) {
    switch choice {
    case "y":
        e = e.NewProject()
    case "n":
        return e, tea.Quit
    }

    return e, nil
}
