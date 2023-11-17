package main

import (
	"bytes"
	"testing"
)

func TestMain(t *testing.T) {
	t.Parallel()
	t.Run("greets", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Chris")

		got := buffer.String()
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
