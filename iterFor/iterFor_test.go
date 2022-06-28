package iterFor

import "testing"

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("want %q but got %q", want, got)
		}
	}

	t.Run("repeat 'a' 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := ("aaaaa")

		assertCorrectMessage(t, got, want)
	})

}
