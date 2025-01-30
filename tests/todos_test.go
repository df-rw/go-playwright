package tests

import (
	"fmt"
	"log"
	"testing"
)

func TestOne(t *testing.T) {
	BeforeEach(t)

	var err error

	// Taken from /cmd/playwright/do-tests.go
	if _, err = page.Goto("http://localhost:4321"); err != nil {
		log.Fatalf("page.Goto() failed: %v", err)
	}

	titleInput := page.Locator("input[name='new-todo']")
	submitButton := page.Locator("button[name='add-new-todo']")
	todos := page.Locator("ul[id='todo-list'] li")

	for i := 0; i < 5; i++ {
		title := fmt.Sprintf("todo number %d", i)

		if err = titleInput.Fill(title); err != nil {
			log.Fatalf("titleInput.Fill() failed: %v", err)
		}

		if err = submitButton.Click(); err != nil {
			log.Fatalf("submitButton.Click() failed: %v", err)
		}

		if err = page.WaitForLoadState(); err != nil {
			log.Fatalf("page.WaitForLoadState() failed: %v", err)
		}
	}

	count, err := todos.Count()
	if err != nil {
		log.Fatalf("todos.Count() failed: %v", err)
	}
	fmt.Println(count, "elements in the list")
}
