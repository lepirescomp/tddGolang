package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type defaultSleeper struct{}

func (d defaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func printOut(w io.Writer, i any) {
	fmt.Fprintln(w, i)
}

func printAndSleep(w io.Writer, s Sleeper) {
	for i := 3; i > 0; i -= 1 {
		printOut(w, i)
		s.Sleep()
	}
	fmt.Fprint(
		w,
		"Go!",
	)

}

func main() {
	w := os.Stdout
	dSleeper := defaultSleeper{}
	printAndSleep(w, dSleeper)
}
