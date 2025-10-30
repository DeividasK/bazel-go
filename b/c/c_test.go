package c

import "testing"

func TestFoo(t *testing.T) {
	result := Foo()
	if result != "Bar" {
		t.Errorf("Expected Bar, got %s", result)
	}
}

func TestBaz(t *testing.T) {
	result := Baz()
	if result != "Hello Baz" {
		t.Errorf("Expected Hello Baz, got %s", result)
	}
}
