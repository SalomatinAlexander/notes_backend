package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Store struct {
	config *Config
	db     *sql.DB
}

func New(c *Config) *Store {
	return &Store{
		config: c,
	}

}

func (s *Store) Open() error {

	db, err := sql.Open(
		"postgres",
		//"postgres://postgres:07120712aA@v2250198.hosted-by-vdsina.ru:5432/note_db?sslmode=disable",
		//postgres://username:password@localhost/db_name?sslmode=disable
		"host=94.103.92.217 port=5432 user=sasha password=07120712aA dbname=notes sslmode=disable",
		//postgres://sasha:07120712aA@94.103.92.217/notes?sslmode=disable
	)
	print("CONNECT...")
	if err != nil {
		print("CONNECT ERROR:" + fmt.Sprint(err))
		return err
	}
	if err := db.Ping(); err != nil {
		print("CONNECT PING ERROR:" + fmt.Sprint(err))
		return err
	}

	s.db = db
	return nil
}

func (s *Store) Close() error {
	if err := s.db.Close(); err != nil {
		return err
	}
	return nil
}
