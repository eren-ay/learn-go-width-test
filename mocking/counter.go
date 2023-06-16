package mocking

import (
	"fmt"
	"io"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

// takes io.writer like a buffer then sends a string
/*
func Countdown(out io.Writer) {
	fmt.Fprint(out, "3")
}
*/

/*
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
}
*/

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		sleeper.Sleep()
	}

	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	finalWord := "Go!"
	fmt.Fprint(out, finalWord)
}
