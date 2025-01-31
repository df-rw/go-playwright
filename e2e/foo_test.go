package tests

import (
	"fmt"
	"testing"
)

func TestParallel(t *testing.T) {
	for i := 0; i < 5; i++ {
		testName := fmt.Sprintf("test-%d", i)
		t.Run(testName, func(t *testing.T) {
			TestOne(t)
		})
	}
}
