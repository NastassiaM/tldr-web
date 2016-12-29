package db

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/gocql/gocql"
)

var session *gocql.Session

// TLDRs returns a slice of all known tldr-pages.
func TLDRs() []byte {
	resArray := [][]byte{}
	var jsonPage []byte
	iter := session.Query("SELECT JSON * FROM pages").Iter()
	for iter.Scan(&jsonPage) {
		resArray = append(resArray, jsonPage)
	}
	if err := iter.Close(); err != nil {
		panic(err)
	}
	res := bytes.Join(resArray, []byte(`,`))

	// insert '[' to the front
	res = append(res, 0)
	copy(res[1:], res[0:])
	res[0] = byte('[')

	// append ']'
	res = append(res, ']')
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
func FindPage(name string) ([]byte, error) {
	var p []byte
	query := fmt.Sprintf("SELECT JSON * FROM pages WHERE name='%s'", name)
	if err := session.Query(query).Scan(&p); err != nil {
		return nil, errors.New("not found")
	}
	return p, nil
}

// AddPage adds new page to the database or replaces the existing one
// with the given name.
func AddPage(p []byte) error {
	query := fmt.Sprintf("INSERT INTO pages JSON '%s'", p)
	q := session.Query(query).Consistency(gocql.One)
	if err := q.Exec(); err != nil {
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
