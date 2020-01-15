package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	expected := "Hello World..."
	if returnValue := Hello(); returnValue != expected {
		t.Error("Hello() = "+returnValue+" , Needed = '"+expected+"'\n")
	}
}
