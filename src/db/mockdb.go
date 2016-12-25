package db

import "model"

var Tldrs map[string]model.Page

// Give us some seed data
func Init() {
	Tldrs = make(map[string]model.Page)
	AddPage(model.Page{Name: "gcc"})
	AddPage(model.Page{Name: "tar"})
}

func FindPage(name string) model.Page {
	p, ok := Tldrs[name]
	if !ok {
		return model.Page{}
	}
	return p
}

func AddPage(p model.Page) model.Page {
	Tldrs[p.Name] = p
	return p
}

func RemovePage(name string) {
	delete(Tldrs, name)
}
