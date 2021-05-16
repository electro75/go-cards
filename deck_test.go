package main

import "testing"

func TestNewDeck(t *testing.T) { // t ==> test handler
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected dec length of 16 but got %v", len(d))
	}
}
