package help

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)


var (
    header = lipgloss.NewStyle().
        Align(lipgloss.Center).
        Bold(true).
        Underline(true)
    body = lipgloss.NewStyle().
        Align(lipgloss.Left)
)

func (h Help) View() string {
    s := ""
    heading := header.Render("Help\n")

    s += heading
    s += "\n"

    keys := ""

    keys += fmt.Sprintln("1: View All Projects")
    keys += fmt.Sprintln("2: View Current Project")
    keys += fmt.Sprintln("3: Open Finder")
    keys += fmt.Sprintln("4: Open Help")

    s += body.Render(keys)
    return s
}
