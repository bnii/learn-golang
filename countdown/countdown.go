package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWorld = "Go!"
const startCounter = 3

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(duration time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

type DefaultSleeper struct{}

func (s DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := startCounter; i >= 1; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWorld)
}

func main() {
	Countdown(os.Stdout, new(DefaultSleeper))
}
