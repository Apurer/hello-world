package morestrings

import "testing"

func TestReverseRunes(t *testing.T) {
    got := ReverseRunes("!oG ,olleH")
    want := "Hello, Go!"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}