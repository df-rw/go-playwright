package tests

import (
	"fmt"
	"os"
	"testing"
)

func MustGetenv(t *testing.T, name string) string {
	t.Helper()

	o := os.Getenv(name)
	if o == "" {
		t.Fatalf("missing environment variable %s", name)
	}

	return o
}

func TestOne(t *testing.T) {
	BeforeEach(t)

	host := MustGetenv(t, "TESTHOST")

	var err error

	// Taken from /cmd/playwright/do-tests.go
	if _, err = page.Goto(host); err != nil {
		t.Fatalf("page.Goto() failed: %v", err)
	}

	titleInput := page.Locator("input[name='new-todo']")
	submitButton := page.Locator("button[name='add-new-todo']")
	todos := page.Locator("ul[id='todo-list'] li")

	for i := 0; i < 5; i++ {
		title := fmt.Sprintf("todo number %d", i)

		if err = titleInput.Fill(title); err != nil {
			t.Fatalf("titleInput.Fill() failed: %v", err)
		}

		if err = submitButton.Click(); err != nil {
			t.Fatalf("submitButton.Click() failed: %v", err)
		}

		if err = page.WaitForLoadState(); err != nil {
			t.Fatalf("page.WaitForLoadState() failed: %v", err)
		}
	}

	count, err := todos.Count()
	if err != nil {
		t.Fatalf("todos.Count() failed: %v", err)
	}
	fmt.Println(count, "elements in the list")
}
