package explorer

type Explorer struct {
    name string // temp variable name
}

func NewExplorer() Explorer {
    return Explorer{}
}

func NewExplorerByName(name string) Explorer {
    return Explorer{
        name: name,
    }
}
