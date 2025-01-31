package main

import (
	"fmt"
	"log"

	"github.com/playwright-community/playwright-go"
)

func main() {
	// Browser install.
	err := playwright.Install()
	if err != nil {
		log.Fatalf("playwright.Install() failed: %v", err)
	}

	// Startup playwright.
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("playwright.Run() failed: %v", err)
	}

	// Launch browser.
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	if err != nil {
		log.Fatalf("pw.Chromium.Launch() failed: %v", err)
	}

	// run tests
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("browser.NewPage() failed: %v", err)
	}

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

	// Cleanup.
	// TODO defer these
	if err = browser.Close(); err != nil {
		log.Fatalf("browser.Close() failed: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("pw.Stop() failed: %v", err)
	}
}
