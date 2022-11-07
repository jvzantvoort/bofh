package main

import (
	"testing"
)

func TestGetExcuse(t *testing.T) {
	message := GetExcuse()
	if len(message) == 0 {
		t.Errorf("got %q, wanted nothing", message)
	}
}

