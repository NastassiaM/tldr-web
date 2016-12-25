package model

// Example represents a single example of command line tool usage
// along with some short description.
type Example struct {
	ID          uint   `json:"id"`
	CmdLine     string `json:"command_line"`
	Description string `json:"desc"`
}

// Page represents a single tl;dr-page for a command line tool.
type Page struct {
	Name        string    `json:"name"`
	Description string    `json:"desc"`
	Examples    []Example `json:"examples"`
}

// Pages is just a collection of Page.
type Pages []Page
