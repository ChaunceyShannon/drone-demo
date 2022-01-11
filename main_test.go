package main

import (
	"testing"
)

func TestPlus(t *testing.T) {
	if plus(1, 1) != 2 {
		t.Error("Error on plus function")
	}
}
