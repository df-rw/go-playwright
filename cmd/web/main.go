package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/go-chi/chi/v5"
)

type Application struct {
	tpl *template.Template
}

func New() *Application {
	tpl := template.Must(template.ParseGlob("templates/*.tmpl"))

	return &Application{
		tpl,
	}
}

func (app *Application) serverError(w http.ResponseWriter, message string, err error) {
	log.Printf("%s (%v)\n", message, err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) render(w http.ResponseWriter, name string, data map[string]any, statusCode int) {
	buf := &bytes.Buffer{}
	err := app.tpl.ExecuteTemplate(buf, name, data)
	if err != nil {
		app.serverError(w, fmt.Sprintf("failed to execute template %s", name), err)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(statusCode)
	_, err = w.Write(buf.Bytes())
	if err != nil {
		app.serverError(w, "Write() failed", err)
	}
}

func main() {
	port := flag.Int("p", 4321, "webserver port")

	app := New()
	r := chi.NewRouter()

	r.Get("/", app.todoHome)               // show all todos
	r.Post("/todo/add", app.todoAdd)       // add a new todo
	r.Post("/todo/update", app.todoUpdate) // update todos on the page

	log.Printf("listening on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), r))
}
