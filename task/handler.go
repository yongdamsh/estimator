package task

import (
	"html/template"
	"log"
	"net/http"
)

type handler struct{}

type model struct {
	Tasks    []Task
	Features []Feature
}

func (h *handler) renderTasks(w http.ResponseWriter, req *http.Request) {
	tasks, err := readTasks()

	if err != nil {
		log.Fatal(err)
	}

	t, err := template.New("tasks").Parse(ListTemplate)

	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, model{tasks, nil})
}

func (h *handler) renderNewTask(w http.ResponseWriter, req *http.Request) {
	t, err := template.New("newTask").Parse(NewTaskTemplate)

	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, nil)
}

func (h *handler) createTask(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	newTask := Task{
		Name:    req.FormValue("name"),
		OrigEst: req.FormValue("estimatedtime"),
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

	http.Redirect(w, req, "/tasks/", http.StatusFound)
}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		if req.URL.Path == "/tasks/new" {
			h.renderNewTask(w, req)
		} else {
			h.renderTasks(w, req)
		}
	case http.MethodPost:
		h.createTask(w, req)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
