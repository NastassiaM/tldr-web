package tldr

type Example struct {
	ID          uint
	CmdLine     string
	Description string
}

type Page struct {
	Name        string
	Description string
	Examples    []Example
}

type Pages []Page
