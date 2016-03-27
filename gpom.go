package main

import (
	"flag"
	"fmt"
	"os/exec"
	"time"
)

const Duration time.Duration = time.Duration(25) * time.Minute

type Timer struct {
	Remaining time.Duration
	Ticker    *time.Ticker
}

var timer Timer

func main() {
	var command string

	flag.Usage = func() {
		fmt.Println("Usage: gpom COMMAND\n")
		fmt.Println("A Pomodoro timer.\n")
		fmt.Println("Author: \n\tBrandon Rice <brandon@blrice.net>\n")
		fmt.Println("Commands:\n\tstart\tStart the timer")
	}

	flag.Parse()

	if len(flag.Args()) > 0 {
		command = flag.Args()[0]
	}

	switch command {
	case "start":
		Start()
	default:
		flag.Usage()
	}
}

func Start() {
	timer = Timer{Duration, time.NewTicker(time.Second)}

	fmt.Println("Starting timer...")

	go func() {
		for {
			select {
			case <-timer.Ticker.C:
				timer.Remaining -= 1 * time.Second
				fmt.Printf("   %s remaining...     \r", timer.Remaining)
			}
		}
	}()

	select {
	case <-time.After(timer.Remaining):
		timer.Ticker.Stop()
		fmt.Println("\nTime's up!\a")
		command := exec.Command("notify-send", "Pomodoro", "Take a break!")
		command.Run()
	}
}
