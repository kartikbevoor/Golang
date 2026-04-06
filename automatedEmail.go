package main

import (
	"log"

	"github.com/robfig/cron/v3"
)

func automatedEmail() {
	// creates a new cron scheduler
	c := cron.New()

	// Run monthly (1st day at 9 AM)
	// Refer the below commented
	c.AddFunc("0 9 1 * *", func() {
		log.Println("Running monthly job...")
		runMonthlyJob()
	})

	// starts the cron scheduler in the background
	c.Start()

	select {} // keep app running
}

func runMonthlyJob() {
	// Fetch customers
	// Calculate balance
	// Send SES email
}

// ┌──────── minute (0–59)
// │ ┌────── hour (0–23)
// │ │ ┌──── day of month (1–31)
// │ │ │ ┌── month (1–12 or *)
// │ │ │ │ ┌ day of week (0–6 or *)
// │ │ │ │ │
// 0 9 1 * *

// ✅ Meaning:
// Minute: 0
// Hour: 9 AM
// Day of month: 1st
// Month: Every month
// Day of week: Any
