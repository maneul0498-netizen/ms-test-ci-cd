package main

import "testing"

func TestMessa(t *testing.T) {
	s := Message("manuel")

	if s != "notification sent to manuel" {
		t.Fatal()
	}

}
