package c

import "testing"

func TestFoo(t *testing.T) {
	result := Foo()
	if result != "Bar" {
		t.Errorf("Expected Bar, got %s", result)
	}
}
