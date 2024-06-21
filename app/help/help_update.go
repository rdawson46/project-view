package help

import tea "github.com/charmbracelet/bubbletea"

func (h Help) Update(tea.Msg) (tea.Model, tea.Cmd) {
    return h, nil
}
