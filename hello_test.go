package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		// %q is a verb that quotes strings
		// see: https://pkg.go.dev/fmt#hdr-Printing
		t.Errorf("got %q want %q", got, want)
	}
}
