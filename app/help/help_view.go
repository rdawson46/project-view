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

    keys += fmt.Sprintln("e: View File in Project")
    keys += fmt.Sprintln("p: View Current Project")
    keys += fmt.Sprintln("f: Open Finder")
    keys += fmt.Sprintln("?: Open Help")
    keys += fmt.Sprintln("q: Quit")

    s += body.Render(keys)
    return s
}
