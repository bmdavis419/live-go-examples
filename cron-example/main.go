package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	// create the scheduler
	loc, err := time.LoadLocation("EST")
	if err != nil {
		panic(err)
	}
	s := gocron.NewScheduler(loc)

	s.Every(1).Second().Do(task)
	s.Every(2).Seconds().Do(func() {
		fmt.Println("Every 2 seconds")
	})

	s.StartBlocking()
}

func task() {
	println("Hello World!")
}
