package maps

import "testing"

func TestSearch(t *testing.T) {
	// "test" like a index number like a tag
	dictionary := map[string]string{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got := Search(dictionary, "test")
		want := "this is just a test"

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, "test")
		}
	})

}
