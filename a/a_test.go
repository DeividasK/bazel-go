package a

import "testing"

func TestHello(t *testing.T) {
	result := Hello()
	expected := "Address: 0x1234567890123456789012345678901234567890; Hello"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
