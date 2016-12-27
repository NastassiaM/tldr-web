package db

import (
	"encoding/json"
	"errors"
	"model"

	"fmt"

	"github.com/gocql/gocql"
)

var session *gocql.Session

// TLDRs returns a slice of all known tldr-pages.
func TLDRs() model.Pages {
	res := make(model.Pages, 0)
	var jsonPage []byte
	var page model.Page
	iter := session.Query("SELECT JSON * FROM pages").Iter()
	for iter.Scan(&jsonPage) {
		if err := json.Unmarshal(jsonPage, &page); err != nil {
			panic(err)
		}
		res = append(res, page)
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

	//AddPage(model.Page{Name: "gcc", Description: "Magic!"})
}

// Close closes the database session.
func Close() {
	session.Close()
}

// FindPage returns page by its name or an empty page and an error
// if nothing is found.
func FindPage(name string) (model.Page, error) {
	var p []byte
	var page model.Page
	query := fmt.Sprintf("SELECT JSON * FROM pages WHERE name='%s'", name)
	if err := session.Query(query).Scan(&p); err != nil {
		return model.Page{}, errors.New("not found")
	}
	if err := json.Unmarshal(p, &page); err != nil {
		return model.Page{}, errors.New("can't decode entry")
	}
	return page, nil
}

// AddPage adds new page to the database or replaces the existing one
// with the given name.
func AddPage(p model.Page) error {
	jsonPage, err := json.Marshal(p)
	if err != nil {
		return err
	}
	query := fmt.Sprintf("INSERT INTO pages JSON '%s'", jsonPage)
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
