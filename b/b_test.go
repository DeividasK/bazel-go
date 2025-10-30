package b

import "testing"

func TestWorld(t *testing.T) {
	result := World()
	if result != "Bar World" {
		t.Errorf("Expected Bar World, got %s", result)
	}
}
