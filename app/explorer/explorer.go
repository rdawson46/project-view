package explorer

import (
	"fmt"
	"os"
    "path"
)

type Explorer struct {
    name  string // temp variable name, copies app.path and appends
    exist bool   // flag to indicate if show prompt user for new project
    localPath string
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

func (e Explorer) NewProject() Explorer {
    newDir := path.Join(e.localPath, e.name)

    err := os.Mkdir(newDir, os.ModePerm)

    if err != nil {
        fmt.Println("Error when creating project dir:", err.Error())
        os.Exit(1)
    }

    e.exist = true

    return e
}
