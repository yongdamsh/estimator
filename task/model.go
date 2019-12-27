package task

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Model struct {
	db *sql.DB
}

func (m *Model) initSchema() error {
	sqlStmt := `
		create table if not exists task (
			id integer not null primary key autoincrement, 
			name text not null,
			estimated text not null,
			elapsed text default '0m',
			createdAt datetime default current_time,
			featureId integer not null references feature on update cascade
		);
		create table if not exists feature (
			id integer not null primary key autoincrement,
			name text not null
		);
	`
	_, err := m.db.Exec(sqlStmt)

	return err
}

func NewModel() *Model {
	db, err := sql.Open("sqlite3", "./data/task.db")

	if err != nil {
		log.Fatal(err)
	}

	m := &Model{db}

	if err := m.initSchema(); err != nil {
		log.Fatal(err)
	}

	return m
}
