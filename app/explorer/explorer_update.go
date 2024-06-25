package explorer

import (
	"fmt"
	"os"
	"path"

	tea "github.com/charmbracelet/bubbletea"
)

func (e Explorer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch message := msg.(type) {

    case tea.WindowSizeMsg:
        e.width = message.Width - 4
        e.height = message.Height - 4
        return e, nil

    case readFileMsg:
        e.content.text = message.content
        return e, nil

    case refreshFileMsg:
        // TODO: read file and update display
        file := message.file
        
        // read correct file
        switch e.focus {
        case todo:
            p := path.Join(e.localPath, e.name, file, "todo.md")
            return e, readFile(p)

        case notes:
            p := path.Join(e.localPath, e.name, file, "notes.md")
            return e, readFile(p)

        default:
            // something bad happened
        }

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

            case "n":
                switch e.focus {
                case notes:
                    e.focus = todo
                case todo:
                    e.focus = notes
                }
                return e, refreshFile(e.tree.entries[e.tree.selected])

            case "enter":
                e.tree.selected = e.tree.current
                return e, refreshFile(e.tree.entries[e.tree.selected])
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

type refreshFileMsg struct{
    file string
}

func refreshFile(file string) tea.Cmd {
    return func() tea.Msg {
        return refreshFileMsg{
            file: file,
        }
    }
}

type readFileMsg struct {
    content string
}

func readFile(p string) tea.Cmd {
    return func() tea.Msg {
        b, err := os.ReadFile(p)

        if err != nil {
            fmt.Println("Could not read file:", err.Error())
            os.Exit(1)
        }

        return readFileMsg {
            content: string(b),
        }
    }
}
