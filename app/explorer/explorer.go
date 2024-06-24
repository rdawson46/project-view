package explorer

import (
	"fmt"
	"os"
	"path"
    "strings"
)

type Explorer struct {
    name  string // temp variable name, copies app.path and appends
    exist bool   // flag to indicate if show prompt user for new project
    localPath string
    tree *TreeData
}

func NewExplorer(p string) Explorer {
    // get current folder location
    wd, err := os.Getwd()

    if err != nil {
        fmt.Println("Could not get working dir")
        os.Exit(1)
    }

    // remove all, but last dir name
    _, projectName := path.Split(wd)

    // check if new
    pojectPath := path.Join(p, projectName)

    info, err := os.Stat(pojectPath)

    var exist bool
    if os.IsNotExist(err) {
        exist = false
        
    } else if !info.IsDir() {
        fmt.Println("Project is not saved as dir")
        os.Exit(1)

    } else if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    } else {
        exist = true
    }

    // TODO: do this in UI
    // if so prompt user if they want to make a new project

    // TODO: figure out how to handle this
    // if not, send to finder or exit
    // figure out how to make a new one

    // open up
    return Explorer{
        name: projectName,
        exist: exist,
        localPath: p,
    }
}

func NewExplorerByName(p, name string) Explorer {
    return Explorer{
        name: name,
    }
}

// TODO: make a tree system
// make tree for project and store in path.Join(e.localPath, e.name)
func (e Explorer) NewProject() Explorer {
    newDir := path.Join(e.localPath, e.name)

    err := os.Mkdir(newDir, os.ModePerm)

    if err != nil {
        fmt.Println("Error when creating project dir:", err.Error())
        os.Exit(1)
    }

    e.exist = true


    // TODO: make tree
    wd, err := os.Getwd()

    if err != nil {
        fmt.Println("Couldn't get wd:", err.Error())
        os.Exit(1)
    }

    entires, err := os.ReadDir(wd)

    if err != nil {
        fmt.Println("Couldn't read wd:", err.Error())
        os.Exit(1)
    }

    for _, entry := range entires {
        if entry.IsDir() {
            // not sure how I want to impl this
        } else {
            name := entry.Name()

            split := strings.Split(name, ".")

            if split[0] == "" {
                continue
            }

            name = split[0]

            /*
            IDEAS:
            1. Make a dir for each file and then make a todo.md and notes.md in the dir ⭐
            2. make sub files
                * [name]-todo.md
                * [name]-notes.md
            */
        }
    }

    return e
}

type TreeData struct {
    entries []string
    current int
}
