package task

import (
	"database/sql"
	"errors"
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
		createdAt datetime default current_timestamp,
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

func (m *Model) tasks() (tasks []Task, err error) {
	q := `
	select task.id, task.name, task.estimated, task.elapsed, task.createdAt, task.featureId, feature.name
	from task join feature on task.featureId = feature.id;
	`

	rows, err := m.db.Query(q)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		task := Task{}

		if err = rows.Scan(
			&task.Id,
			&task.Name,
			&task.Estimated,
			&task.Elapsed,
			&task.CreatedAt,
			&task.Feature.Id,
			&task.Feature.Name,
		); err != nil {
			return
		}

		tasks = append(tasks, task)
	}

	return
}

func (m *Model) features() (features []Feature, err error) {
	q := `
	select * from feature;
	`

	rows, err := m.db.Query(q)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		feature := Feature{}

		if err = rows.Scan(
			&feature.Id,
			&feature.Name,
		); err != nil {
			return
		}

		features = append(features, feature)
	}

	return
}

func (m *Model) addTask(task *Task) (err error) {
	sqlStmt := `
	insert into task (name, estimated, featureId) values (?, ?, ?);
	`

	result, err := m.db.Exec(sqlStmt, task.Name, task.Estimated, task.Feature.Id)

	numRows, err := result.RowsAffected()

	if err != nil {
		return
	}

	if numRows != 1 {
		return errors.New("task did not added")
	}

	return
}

func (m *Model) updateTask(task *Task) (err error) {
	sqlStmt := `
	update task set name = ?, estimated = ?, elapsed = ?, featureId = ? where id = ?;
	`

	result, err := m.db.Exec(sqlStmt,
		task.Name,
		task.Estimated,
		task.Elapsed,
		task.Feature.Id,
		task.Id,
	)

	numRows, err := result.RowsAffected()

	if err != nil {
		return
	}

	if numRows != 1 {
		return errors.New("task did not added")
	}

	return
}

func (m *Model) task(id int) (task Task, err error) {
	q := `
	select * from task where id = ?;
	`

	row := m.db.QueryRow(q, id)

	err = row.Scan(
		&task.Id,
		&task.Name,
		&task.Estimated,
		&task.Elapsed,
		&task.CreatedAt,
		&task.Feature.Id,
	)

	return
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
