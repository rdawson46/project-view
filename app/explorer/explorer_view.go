package explorer

import (
    "fmt"
    "github.com/charmbracelet/lipgloss"
)

var (
    dir_holder = lipgloss.NewStyle().
        Align(lipgloss.Left).
        Border(lipgloss.RoundedBorder(), true).
        BorderForeground(lipgloss.Color("#454545")).
        Padding(1, 1)

    selected = lipgloss.NewStyle().
        Bold(true).
        Foreground(lipgloss.Color("45"))

    file_view = lipgloss.NewStyle() 
)

func (e Explorer) View() string {
    if !e.exist {
        return e.newProjectPrompt()
    }

    return e.projectView()
}

func (e Explorer) projectView() string {
    if e.tree != nil {
        var s string

        s += e.dirView()

        return s
    }

    return "No tree found"
}

func (e Explorer) dirView() string {
    dir_holder = dir_holder.Width(e.width / 5).Height(e.height)

    var s string

    for i, entry := range e.tree.entries {
        if i == e.tree.current {
            s += selected.Render(fmt.Sprintf("%s", entry))
            s += "\n"
        } else {
            s += fmt.Sprintln(entry)
        }
    }

    return dir_holder.Render(s)
}

func (e Explorer) textView() string {
    return ""
}


func (e Explorer) newProjectPrompt() string {
    return fmt.Sprintf("Would you like to make a new project (%s)? (y/n)", e.name)
}
