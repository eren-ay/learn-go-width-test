package mocking

import (
	"bytes"
	"testing"
)

/*
func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := wantCount(3)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func wantCount(length int) string {
	countDown := ""

	for i := 3; i > 0; i-- {

		countDown = countDown + fmt.Sprint(i)
		countDown = countDown + "\n"
	}
	return countDown
}
*/

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
	}
}
