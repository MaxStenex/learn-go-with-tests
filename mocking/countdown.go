package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3
const delaySeconds = 1

const sleepOperation = "sleep"
const writeOperation = "write"

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type SpyCountdownOperations struct {
	Calls []string
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (s *ConfigurableSleeper) Sleep() {
	s.sleep(s.duration)
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleepOperation)
}

func (s *SpyCountdownOperations) Write(b []byte) (n int, err error) {
	s.Calls = append(s.Calls, writeOperation)

	return
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{delaySeconds * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
