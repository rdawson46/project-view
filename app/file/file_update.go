package file

import tea "github.com/charmbracelet/bubbletea"

func (f File) Update(tea.Msg) (tea.Model, tea.Cmd) {
    return f, nil
}
