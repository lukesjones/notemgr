package main

import "testing"

func TestDeleteList(t *testing.T) {
	expected := "the delete function is not yet in use!"

	output := deleteList()

	if output != expected {
		t.Error("Output is different to expected!")
	}
}
