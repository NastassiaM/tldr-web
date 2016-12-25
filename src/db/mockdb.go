package db

import (
	"errors"
	"model"
)

var tldrs map[string]model.Page

// OldTLDRs returns a slice of all known tldr-pages.
func OldTLDRs() model.Pages {
	res := make([]model.Page, 0, len(tldrs))
	for _, v := range tldrs {
		res = append(res, v)
	}

	return res
}

// OldInit is here to initialize our database and add some seed data.
func OldInit() {
	tldrs = make(map[string]model.Page)
	AddPage(model.Page{Name: "gcc"})
	AddPage(model.Page{Name: "tar"})
}

// OldFindPage returns page by its name or an empty page and an error
// if nothing is found.
func OldFindPage(name string) (model.Page, error) {
	p, ok := tldrs[name]
	if !ok {
		return model.Page{}, errors.New("not found")
	}
	return p, nil
}

// OldAddPage adds new page to the database or replaces the existing one
// with the given name.
func OldAddPage(p model.Page) error {
	tldrs[p.Name] = p
	return nil
}

// OldRemovePage removes entry from database.
func OldRemovePage(name string) error {
	delete(tldrs, name)
	return nil
}
