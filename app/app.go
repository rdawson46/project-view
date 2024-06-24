package app

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rdawson46/project-view/app/explorer"
	"github.com/rdawson46/project-view/app/file"
	"github.com/rdawson46/project-view/app/finder"
	"github.com/rdawson46/project-view/app/help"
)

const (
    Explorer = iota + 1// search for files in project base
    File    // notepad and to do
    Finder  // screen for changing projects
    Help
)


type App struct {
   currentScreen int 
   screen        tea.Model
   height, width int
   path          string // path ~/.local/share/project-view
}

func NewApp() App {
    // find directory in ~/.local/share/project-view after project
    home, err := os.UserHomeDir()

    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    path := fmt.Sprintf("%s/.local/share/project-view", home)

    info, err := os.Stat(path)

    if os.IsNotExist(err) {
        err = os.MkdirAll(path, os.ModePerm)

        if err != nil {
            fmt.Println("Could not make directory")
            os.Exit(1)
        }


        info, err = os.Stat(path)

        if err != nil {
            fmt.Println("Broke at line 54")
            os.Exit(1)
        }
    }

    if !info.IsDir() {
        fmt.Println("path is file or doesn't exist")
        os.Exit(1)
    }

    // if not found, ask if one should be created (explorer job?)
    return App {
        currentScreen: 0,
        screen: explorer.NewExplorer(path, 0, 0),
        path: path,
    }
}

// simple idea, might be too simple
func (a App) changeScreen(screenType int, name ...string) App {
    switch screenType {
    case Explorer:
        a.currentScreen = screenType
        a.screen = explorer.NewExplorer(a.path, a.height, a.width)

    case File:
        a.currentScreen = screenType
        a.screen = file.NewFile()

    case Finder:
        a.currentScreen = screenType
        a.screen = finder.NewFinder()

    case Help:
        a.currentScreen = screenType
        a.screen = help.NewHelp()
    }

    return a
}

func (a App) resize(height, width int) App {
    a.height = height
    a.width = width
    return a
}
