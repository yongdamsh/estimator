package task

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Handler struct {
	model *Model
}

func (h *Handler) renderTasks(w http.ResponseWriter, req *http.Request) {
	tasks, err := h.model.tasks()

	if err != nil {
		log.Fatal(err)
	}

	t, err := template.New("tasks").Parse(ListTemplate)

	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, tasks)
}

func (h *Handler) renderNewTask(w http.ResponseWriter, req *http.Request) {
	var wg sync.WaitGroup
	f := make(chan []Feature)

	wg.Add(1)

	go func() {
		defer wg.Done()

		features, err := h.model.features()

		if err != nil {
			log.Fatal(err)
		}

		f <- features
	}()

	t, err := template.New("newTask").Parse(NewTaskTemplate)

	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, <-f)
	wg.Wait()
}

func (h *Handler) createTask(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	featureId, err := strconv.Atoi(req.FormValue("featureId"))

	if err != nil {
		log.Fatal(err)
	}

	newTask := Task{
		Feature: Feature{
			Id: featureId,
		},
		Name:      req.FormValue("name"),
		Estimated: req.FormValue("estimated"),
	}

	err = h.model.addTask(&newTask)

	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, req, "/tasks/", http.StatusFound)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
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
