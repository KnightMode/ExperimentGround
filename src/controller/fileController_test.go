package controller

import (
	"testing"
)

func TestRouter(t *testing.T) {
	expected := "Hello, world!"
	actual := "Hello, world!"
	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}
