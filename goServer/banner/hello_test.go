package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	expected := "\n"
	expected += "            ____                           "
	expected += "  __ _  ___/ ___|  ___ _ ____   _____ _ __ "
	expected += " / _` |/ _ \___ \ / _ \ '__\ \ / / _ \ '__|"
	expected += "| (_| | (_) |__) |  __/ |   \ V /  __/ |   "
	expected += " \__, |\___/____/ \___|_|    \_/ \___|_|   "
	expected += " |___/\n"
	if returnValue := Hello(); returnValue != expected {
		t.Error("Hello() = "+returnValue+" , Needed = '"+expected+"'\n")
	}
}
