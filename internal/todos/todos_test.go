package todos

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAdd(t *testing.T) {
	todos := New()

	_, err := todos.Add("one", false)
	if err != nil {
		t.Fatalf("todos.Add(): got unexpected error: %v", err)
	}

	numTodos := len(todos.All())
	if numTodos != 1 {
		t.Errorf("todos.All(): wrong number of todos; got %d, want %d\n", numTodos, 1)
	}
}

func TestAll(t *testing.T) {
	todos := New()
	want := []Todo{
		{"one", false},
		{"two", false},
		{"three", false},
		{"four", false},
		{"five", false},
	}
	for _, v := range want {
		todos.Add(v.Title, v.Done)
	}

	got := todos.All()
	if diff := cmp.Diff(want, got); diff != "" { // note order here
		t.Errorf("todos.All() mismatch (-want, +got):\n%s", diff)
	}
}

func TestSetAllDone(t *testing.T) {
	todos := New()
	want := []Todo{
		{"one", false},
		{"two", false},
		{"three", false},
		{"four", false},
		{"five", false},
	}
	for _, v := range want {
		todos.Add(v.Title, v.Done)
	}

	// All should be false, as per setup.
	err := todos.SetAllDone(false)
	if err != nil {
		t.Fatalf("todos.SetAllDone(): got unexpected error: %v", err)
	}
	got := todos.All()

	if diff := cmp.Diff(want, got); diff != "" { // note order here
		t.Fatalf("todos.All() mismatch (-want, +got):\n%s", diff)
	}

	// All should be true.
	err = todos.SetAllDone(true)
	if err != nil {
		t.Fatalf("todos.SetAllDone(): got unexpected error: %v", err)
	}
	got = todos.All()

	for i := range want {
		want[i].Done = true
	}

	if diff := cmp.Diff(want, got); diff != "" { // note order here
		t.Fatalf("todos.All() mismatch (-want, +got):\n%s", diff)
	}
}

func TestSetDone(t *testing.T) {
	todos := New()
	id, err := todos.Add("one", false)
	if err != nil {
		t.Fatalf("todos.Add(): got unexpected error: %v", err)
	}

	want := true
	err = todos.SetDone(id, want)
	if err != nil {
		t.Fatalf("todos.SetDone(): got unexpected error: %v", err)
	}

	got, err := todos.IsDone(id)
	if err != nil {
		t.Fatalf("todos.SetDone(): got unexpected error: %v", err)
	}

	if want != got {
		t.Errorf("todos.IsDone(): got %v, want %v", got, want)
	}
}

func TestToggleDone(t *testing.T) {
	todos := New()
	id, err := todos.Add("one", false)
	if err != nil {
		t.Fatalf("todos.Add(): got unexpected error: %v", err)
	}

	err = todos.ToggleDone(id)
	if err != nil {
		t.Fatalf("todos.ToggleDone(): got unexpected error: %v", err)
	}

	want := true
	got, err := todos.IsDone(id)
	if err != nil {
		t.Fatalf("todos.IsDone(): got unexpected error: %v", err)
	}

	if got != want {
		t.Errorf("todos.IsDone(): got %v, want %v", got, want)
	}
}
