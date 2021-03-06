package main

import (
	"fmt"
	"time"

	"github.com/andybrewer/mack"
)

func main() {
	// Tickers use a similar mechanism to timers: a
	// channel that is sent values. Here we'll use the
	// `select` builtin on the channel to await the
	// values as they arrive.
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Tickers can be stopped like timers. Once a ticker
	// is stopped it won't receive any more values on its
	// channel.
	time.Sleep(5 * time.Second)

	dialog := mack.DialogOptions{
		Text:    "Time's up",    // Required
		Buttons: "Stop, Snooze", // Optional - Comma separated list, max of 3
	}
	response, err := mack.DialogBox(dialog)

	if err != nil {
		panic(err)
	}

	if response.Clicked == "Stop" {
		fmt.Println("Stop")
		ticker.Stop()
	} else {
		fmt.Println("Snoozing for 3 seconds")
		time.Sleep(3 * time.Second)
	}

	done <- true
	fmt.Println("Ticker stopped")
}
