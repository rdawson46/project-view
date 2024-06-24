package explorer

import "fmt"

func (e Explorer) View() string {
    if !e.exist {
        return e.newProjectPrompt()
    }

    return e.projectView()
}

func (e Explorer) projectView() string {
    return "project exists"
}

func (e Explorer) newProjectPrompt() string {
    return fmt.Sprintf("Would you like to make a new project (%s)? (y/n)", e.name)
}


