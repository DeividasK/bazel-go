package a

import "testing"

func TestHello(t *testing.T) {
	result := Hello()
	if result != "Hello" {
		t.Errorf("Expected Hello, got %s", result)
	}
}
