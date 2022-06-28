package go_hello_world

import "testing"

func TestGreeter(t *testing.T) {
	expected := "hello world!"
	actual := HelloWorld()
	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}
