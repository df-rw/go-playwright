package main

import (
	"net/http"
	"regexp"
	"strconv"
)

func (app *Application) todoHome(w http.ResponseWriter, r *http.Request) {
	pageData := map[string]any{
		"Todos": app.todos.All(),
	}

	app.render(w, "todos", pageData, http.StatusOK)
}

func (app *Application) todoAdd(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	title := r.FormValue("new-todo")

	_, _ = app.todos.Add(title, false)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *Application) todoUpdate(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	app.todos.SetAllDone(false)

	re := regexp.MustCompile(`todo-(\d*)`)
	for k, _ := range r.Form {
		matches := re.FindStringSubmatch(k)
		if len(matches) == 2 {
			i, _ := strconv.Atoi(matches[1])
			app.todos.SetDone(i, true)
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
