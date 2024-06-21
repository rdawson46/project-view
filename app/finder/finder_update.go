package finder

import tea "github.com/charmbracelet/bubbletea"

func (f Finder) Update(tea.Msg) (tea.Model, tea.Cmd) {
    return f, nil
}
