package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

// Cron = time-based job scheduler
// Cron in Go is used to schedule tasks to run automatically at specific times or intervals — just like Linux
// cron jobs, but inside your Go application.

// basic function
func basicCron() {
	// creates a new cron scheduler instance.
	c := cron.New()

	// Prevent Overlapping
	// c := cron.New(
	// 	cron.WithChain(
	// 		cron.SkipIfStillRunning(cron.DefaultLogger),
	// 	),
	// )

	// cron.WithChain(...):
	// lets you define a chain of job wrappers (middleware).
	// These wrappers modify how jobs behave when they run.
	// Think of it like adding rules around execution.

	// cron.SkipIfStillRunning(cron.DefaultLogger)
	// If a scheduled job is still running when the next trigger happens: Skip the new run entirely

	// cron.DefaultLogger: This is the logger used by the cron package.

	// DelayIfStillRunning:
	// Instead of skipping, it waits until the previous job finishes
	// Then runs the delayed job

	// "@every 1m" → every minute
	// "0 0 * * *" → every day at midnight
	c.AddFunc("@every 10s", func() {
		fmt.Println("Running every 10 seconds:", time.Now())
	})

	// starts the cron scheduler in its own goroutine
	c.Start()
	select {} // keep program running

}

// Cron Expression Format
// * * * * *
// | | | | |
// | | | | └── Day of week (0-6)
// | | | └──── Month (1-12)
// | | └────── Day of month (1-31)
// | └──────── Hour (0-23)
// └────────── Minute (0-59)

// Examples
// | Cron Expression | Meaning                        |
// | --------------- | ------------------------------ |
// | `0 9 * * *`     | Every day at 9:00 AM           |
// | `*/5 * * * *`   | Every 5 minutes                |
// | `0 0 * * 0`     | Every Sunday at midnight       |
// | `@every 10s`    | Every 10 seconds (Go shortcut) |
// | `@daily`        | Once per day                   |

// Cron to use seconds field:
// c := cron.New(cron.WithSeconds())
// * * * * * *
// | | | | | |
// | | | | | └── Day of week
// | | | | └──── Month
// | | | └────── Day
// | | └──────── Hour
// | └────────── Minute
// └──────────── Second

// Cron job without ananymous function (named function)
func tellTime() {
	fmt.Println("Running every 10 seconds:", time.Now())
}

func CronWithNamedFunction() {
	c := cron.New()

	c.AddFunc("@every 10s", tellTime)

	c.Start()

}
