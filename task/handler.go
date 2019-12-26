package task

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type handler struct{}

func (h *handler) retrieveTasks(w http.ResponseWriter, req *http.Request) {
	tasks, err := readTasks()

	if err != nil {
		log.Fatal(err)
	}

	t, err := template.New("tasks").Parse(Template)

	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, tasks)
}

func (h *handler) createTask(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Fatal(err)
	}

	var newTask Task

	err = json.Unmarshal(body, &newTask)

	if err != nil {
		log.Fatal(err)
	}

	tasks, err := readTasks()

	if err != nil {
		log.Fatal(err)
	}

	tasks = append(tasks, newTask)
	err = saveTasks(tasks)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Location", "/tasks")
	w.WriteHeader(http.StatusCreated)
}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		h.retrieveTasks(w, req)
	case http.MethodPost:
		h.createTask(w, req)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
