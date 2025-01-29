package main

import (
	"net/http"
	"regexp"
	"strconv"
)

type Todo struct {
	Title string
	Done  bool
}

var todos []Todo

func (app *Application) todoHome(w http.ResponseWriter, r *http.Request) {
	pageData := map[string]any{
		"Todos": todos,
	}

	app.render(w, "todos", pageData, http.StatusOK)
}

func (app *Application) todoAdd(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("new-todo")

	if name != "" {
		newTodo := Todo{
			Title: name,
			Done:  false,
		}

		todos = append(todos, newTodo)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *Application) todoUpdate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// We only get POSTed stuff that is enabled. Turn all todos off
	// and then we'll enable just the stuff that we receive.

	for i := 0; i < len(todos); i++ {
		todos[i].Done = false
	}

	re := regexp.MustCompile(`todo-(\d*)`)
	for k, _ := range r.Form {
		matches := re.FindStringSubmatch(k)
		if len(matches) == 2 {
			i, _ := strconv.Atoi(matches[1])
			todos[i].Done = true
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
