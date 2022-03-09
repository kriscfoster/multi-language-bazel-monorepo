package greeter

import "testing"

func TestGreeter(t *testing.T) {
	expected := "Hello World!"
	greeting := Greet()
	if greeting != expected {
		t.Errorf("expected %q but got %q", expected, greeting)
	}
}
