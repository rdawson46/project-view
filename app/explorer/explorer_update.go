package explorer

import tea "github.com/charmbracelet/bubbletea"

func (e Explorer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch message := msg.(type) {

    case tea.WindowSizeMsg:
        e.width = message.Width - 5
        e.height = message.Height - 5
        return e, nil

    case tea.KeyMsg:
        switch e.exist {
        case false:
            return e.newPromptUpdate(message.String())

        default:
            switch message.String(){
            case "j":
                if e.tree != nil && e.tree.current < len(e.tree.entries) - 1 {
                    e.tree.current++
                }
            case "k":
                if e.tree != nil && e.tree.current > 0 {
                    e.tree.current--
                }
            }
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
