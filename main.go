package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {
	var interval int
	var sink string

	flag.IntVar(&interval, "interval", 5, "Interval (in secionds")
	flag.StringVar(&sink, "sink", "http://localhost", "Sink to send events")

	flag.Parse()

	fmt.Printf("Interval: %d\n", interval)
	fmt.Printf("Sink: %s\n", sink)

	timeEvents := make(chan string)

	// Emit an event on the given interval
	go func() {
		intervalDuration := time.Duration(interval) * time.Second
		for {
			t := time.Now()
			timeEvents <- strconv.FormatInt(t.Unix(), 10)
			time.Sleep(intervalDuration)
		}
	}()

	// Deliver the event to the provided sink
	for {
		event := <-timeEvents
		emitEvent(sink, event)
	}
}

func emitEvent(sink string, event string) {
	fmt.Printf("Sending event to %s: %s\n", sink, event)
	_, err := http.Post(sink, "text/plain", bytes.NewBuffer([]byte(event)))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
