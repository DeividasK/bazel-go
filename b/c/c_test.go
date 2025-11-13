package c

import "testing"

func TestFoo(t *testing.T) {
	result := Foo()
	expected := "Bar"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
	t.Logf("Foo result: %s", result)
}

func TestBaz(t *testing.T) {
	result := Baz()

	expected := "Address: 0x1234567890123456789012345678901234567890; Hello Baz"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
