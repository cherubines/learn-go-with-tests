package intergers

import (
	"testing"
)

func TestAAdder(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("test 2+2", func(t *testing.T) {
		got := Add(2, 2)
		want := 4

		assertCorrectMessage(t, got, want)
	})

}
