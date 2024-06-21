package main

import (
    "os"
    "fmt"
    "github.com/rdawson46/project-view/app"
    tea "github.com/charmbracelet/bubbletea"
)

func main() {
    program := tea.NewProgram(app.NewApp(), tea.WithAltScreen())

    if _, err := program.Run(); err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
}
