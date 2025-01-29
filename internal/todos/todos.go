package todos

import "errors"

var (
	ErrInvalidID = errors.New("invalid id")
)

type Todo struct {
	Title string
	Done  bool
}

type Todos struct {
	todos []Todo
}

func New() *Todos {
	return &Todos{}
}

func (t *Todos) All() []Todo {
	return t.todos
}

func (t *Todos) Add(title string, done bool) (int, error) {
	t.todos = append(t.todos, Todo{
		Title: title,
		Done:  done,
	})
	return len(t.todos), nil
}

func (t *Todos) ToggleDone(id int) error {
	if id >= len(t.todos) {
		return ErrInvalidID
	}

	t.todos[id].Done = !t.todos[id].Done

	return nil
}

func (t *Todos) SetDone(id int, done bool) error {
	if id >= len(t.todos) {
		return ErrInvalidID
	}

	t.todos[id].Done = done

	return nil
}

func (t *Todos) SetAllDone(done bool) error {
	for i := range t.todos {
		t.todos[i].Done = done
	}
	return nil
}
