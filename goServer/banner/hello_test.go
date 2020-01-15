package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	expected := "\n"
	expected += "            ____                           \n"
	expected += "  __ _  ___/ ___|  ___ _ ____   _____ _ __ \n"
	expected += " / _` |/ _ \\___ \\ / _ \\ '__\\ \\ / / _ \\ '__|\n"
	expected += "| (_| | (_) |__) |  __/ |   \\ V /  __/ |   \n"
	expected += " \\__, |\\___/____/ \\___|_|    \\_/ \\___|_|   \n"
	expected += " |___/\n"
	if returnValue := Hello(); returnValue != expected {
		t.Error("Hello() = "+returnValue+" , Needed = '"+expected+"'\n")
	}
}
