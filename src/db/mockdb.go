package db

import "model"
import "errors"

var tldrs map[string]model.Page

// TLDRs returns a slice of all known tldr-pages.
func TLDRs() model.Pages {
	res := make([]model.Page, 0, len(tldrs))
	for _, v := range tldrs {
		res = append(res, v)
	}

	return res
}

// Init is here to initialize our database and add some seed data.
func Init() {
	tldrs = make(map[string]model.Page)
	AddPage(model.Page{Name: "gcc"})
	AddPage(model.Page{Name: "tar"})
}

// FindPage returns page by its name or an empty page and an error
// if nothing is found.
func FindPage(name string) (model.Page, error) {
	p, ok := tldrs[name]
	if !ok {
		return model.Page{}, errors.New("not found")
	}
	return p, nil
}

// AddPage adds new page to the database or replaces the existing one
// with the given name.
func AddPage(p model.Page) {
	tldrs[p.Name] = p
}

// RemovePage removes entry from database.
func RemovePage(name string) {
	delete(tldrs, name)
}
