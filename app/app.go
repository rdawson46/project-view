package app

import (
    tea "github.com/charmbracelet/bubbletea"
    exp "github.com/rdawson46/project-view/app/explorer"
    fil "github.com/rdawson46/project-view/app/file"
    fin "github.com/rdawson46/project-view/app/finder"
    hel "github.com/rdawson46/project-view/app/help"
)

const (
    explorer = iota + 1// search for files in project base
    file    // notepad and to do
    finder  // screen for changing projects
    help
)


type App struct {
   currentScreen int 
   screen tea.Model
   height, width int
}

func NewApp() App {
    // find directory in ~/.local/named after project

    // if not found, ask if one should be created (explorer job?)
    return App {
        currentScreen: 0,
        screen: exp.NewExplorer(),
    }
}

// simple idea, might be too simple
func (a App) changeScreen(screenType int, name ...string) App {
    switch screenType {
    case explorer:
        a.currentScreen = screenType
        a.screen = exp.NewExplorer()

    case file:
        a.currentScreen = screenType
        a.screen = fil.NewFile()

    case finder:
        a.currentScreen = screenType
        a.screen = fin.NewFinder()

    case help:
        a.currentScreen = screenType
        a.screen = hel.NewHelp()

    default:

    }

    return a
}
