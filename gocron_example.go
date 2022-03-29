package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func testPrint() {
	fmt.Println("test cron ", time.Now())
}

func TestGocron() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(5).Seconds().Do(testPrint)

	// strings parse to duration
	s.Every("5m").Do(testPrint)

	s.Every(5).Days().Do(testPrint)

	s.Every(1).Month(1, 2, 3).Do(testPrint)

	// set time
	s.Every(1).Day().At("10:30").Do(testPrint)

	// set multiple times
	s.Every(1).Day().At("10:30;08:00").Do(testPrint)

	s.Every(1).Day().At("10:30").At("08:00").Do(testPrint)

	// Schedule each last day of the month
	s.Every(1).MonthLastDay().Do(testPrint)

	// Or each last day of every other month
	s.Every(2).MonthLastDay().Do(testPrint)

	// cron expressions supported
	s.Cron("*/1 * * * *").Do(testPrint) // every minute

	// you can start running the scheduler in two different ways:
	// starts the scheduler asynchronously
	s.StartAsync()
	// starts the scheduler and blocks current execution path
	s.StartBlocking()
}
