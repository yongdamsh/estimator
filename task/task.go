package task

import (
	"encoding/json"
	"io/ioutil"
)

var Handler = new(handler)

func readTasks() (tasks []Task, err error) {
	b, err := ioutil.ReadFile("./data/tasks.json")

	if err != nil {
		return
	}

	err = json.Unmarshal(b, &tasks)

	return
}

func saveTasks(tasks []Task) error {
	b, err := json.MarshalIndent(tasks, "", "\t")

	if err != nil {
		return err
	}

	return ioutil.WriteFile("./data/tasks.json", b, 0644)
}
