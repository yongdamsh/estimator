package task

import (
	"encoding/csv"
	"io"
	"os"
	"time"
)

var Handler = new(handler)

func readTasks() (tasks []Task, err error) {
	f, err := os.Open("./data/task.csv")
	defer f.Close()

	if err != nil {
		return
	}

	r := csv.NewReader(f)
	r.ReuseRecord = true

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		task := Task{
			OrigEst: record[1],
			CurEst:  record[2],
			Elapsed: record[3],
			Feature: Feature{
				Name: record[5],
			},
			Name: record[6],
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func saveTasks(tasks []Task) (err error) {
	f, err := os.OpenFile("./data/task.csv", os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	defer f.Close()

	if err != nil {
		return
	}

	w := csv.NewWriter(f)

	records := make([][]string, len(tasks))

	for index, task := range tasks {
		records[index] = []string{
			"1",
			task.OrigEst,
			task.OrigEst,
			"0m",
			time.Now().String(),
			task.Feature.Name,
			task.Name,
		}
	}

	return w.WriteAll(records)
}
