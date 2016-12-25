package db

import (
	"errors"
	"model"

	"fmt"

	"github.com/gocql/gocql"
)

var session *gocql.Session

// TLDRs returns a slice of all known tldr-pages.
func TLDRs() model.Pages {
	res := make(model.Pages, 0)
	var name, description string
	iter := session.Query("SELECT * FROM pages").Iter()
	for iter.Scan(&name, &description) {
		res = append(res, model.Page{Name: name, Description: description})
	}
	if err := iter.Close(); err != nil {
		panic(err)
	}

	return res
}

// Init is here to initialize our database and add some seed data.
func Init() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "tldr"
	session, _ = cluster.CreateSession()

	AddPage(model.Page{Name: "gcc", Description: "Magic!"})
}

// Close closes the database session.
func Close() {
	session.Close()
}

// FindPage returns page by its name or an empty page and an error
// if nothing is found.
func FindPage(name string) (model.Page, error) {
	var _name, _description string
	query := fmt.Sprintf("SELECT * FROM pages WHERE name='%s'", name)
	if err := session.Query(query).Scan(&_name, &_description); err != nil {
		return model.Page{}, errors.New("not found")
	}
	return model.Page{Name: _name, Description: _description}, nil
}

// AddPage adds new page to the database or replaces the existing one
// with the given name.
func AddPage(p model.Page) error {
	query := fmt.Sprintf("INSERT INTO pages (name, description) VALUES ('%s', '%s')",
		p.Name, p.Description)
	if err := session.Query(query).Exec(); err != nil {
		return err
	}
	return nil
}

// RemovePage removes entry from database.
func RemovePage(name string) error {
	query := fmt.Sprintf("DELETE FROM pages WHERE name = '%s'", name)
	if err := session.Query(query).Exec(); err != nil {
		return err
	}
	return nil
}
