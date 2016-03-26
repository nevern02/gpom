package main

import (
	"flag"
	"fmt"
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

	flag.Parse()

	if len(flag.Args()) > 0 {
		command = flag.Args()[0]
	}

	switch command {
	case "start":
		Start()
	default:
		fmt.Println("Invalid command.")
	}
}

func Start() {
	timer = Timer{Duration, time.NewTicker(time.Second)}
	channel := timer.Ticker.C

	fmt.Println("Starting timer...")

	go func() {
		for {
			select {
			case <-channel:
				timer.Remaining -= 1 * time.Second
				fmt.Printf("   %s remaining...     \r", timer.Remaining)
			}
		}
	}()

	select {
	case <-time.After(timer.Remaining):
		timer.Ticker.Stop()
		fmt.Println("\nTime's up!")
		break
	}
}
