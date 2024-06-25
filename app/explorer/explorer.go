package explorer

import (
	"fmt"
	"os"
	"path"
    "strings"
)

const (
    todo = iota
    notes
)

type Explorer struct {
    name          string // temp variable name, copies app.path and appends
    exist         bool   // flag to indicate if show prompt user for new project
    localPath     string
    focus         int
    tree          *TreeData
    content       *FileContent
    height, width int
}

func NewExplorer(p string, height, width int) Explorer {
    // get current folder location
    wd, err := os.Getwd()

    if err != nil {
        fmt.Println("Could not get working dir")
        os.Exit(1)
    }

    // remove all, but last dir name
    _, projectName := path.Split(wd)

    // check if new
    projectPath := path.Join(p, projectName)

    info, err := os.Stat(projectPath)

    var exist bool
    var tree *TreeData
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

        entries, err := os.ReadDir(projectPath)

        if err != nil {
            fmt.Println("Error when reading known dir:", err.Error())
            os.Exit(1)
        }

        names := make([]string, len(entries))
        for i, entry := range entries {
            names[i] = entry.Name()
        }

        tree = newTreeData(names)

        if tree == nil {
            fmt.Println("Tree wasn't produced")
            os.Exit(1)
        }
    }

    // open up
    return Explorer{
        name: projectName,
        exist: exist,
        localPath: p,
        tree: tree,
        content: &FileContent{},
        focus: notes,
        height: max(height - 2, 0),
        width: max(width - 2, 0),
    }
}

func NewExplorerByName(p, name string) Explorer {
    return Explorer{
        name: name,
    }
}

// TODO: make a tree system
// make tree for project and store in path.Join(e.localPath, e.name)

// BUG: need a way to remove/filter certain files/folder
// ex: node_modules, go.mod, ...
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

    treeEntries := make([]string, 0)

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
            newFilePath := path.Join(newDir, name)
            err := os.Mkdir(newFilePath, os.ModePerm)

            if os.IsExist(err) {
                continue
            }

            if err != nil {
                fmt.Println("Error making dir for file:", err.Error())
                os.Exit(1)
            }


            _, err = os.Create(path.Join(newFilePath, "todo.md"))

            if err != nil {
                fmt.Println("Error when creating todo for:", newFilePath, err.Error())
                os.Exit(1)
            }

            _, err = os.Create(path.Join(newFilePath, "notes.md"))

            if err != nil {
                fmt.Println("Error when creating notes for:", newFilePath, err.Error())
                os.Exit(1)
            }

            treeEntries = append(treeEntries, name)
        }
    }

    e.tree = newTreeData(treeEntries)
    return e
}

type TreeData struct {
    entries  []string // name of each dir
    current  int // idx of currently highlighted/selected dir
    selected int
}

func newTreeData(entries []string) *TreeData {
    return &TreeData{
        entries: entries,
        current: 0,
    }
}

func max(x, y int) int {
    if x > y {
        return x
    } else {
        return y
    }
}

type FileContent struct {
    text string
}
