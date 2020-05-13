// Using a separate morestrings_test package means that I can’t access the unexported fields and functions in my package.
package arithmetics_test

import (
    "testing"
    . "github.com/Apurer/hello-world/arithmetics"
)

func TestSum(t *testing.T) {
    got := Sum(7, 9)
    want := 16

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}