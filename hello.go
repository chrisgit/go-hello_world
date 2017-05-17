package main

import (
	"fmt"
	"time"
)

// Greeting messages
const (
	GoodMorning   = "Good morning"
	GoodAfternoon = "Good afternoon"
	GoodEvening   = "Good evening"
)

func greeting(t time.Time) string {
	switch {
	case t.Hour() < 12:
		return GoodMorning
	case t.Hour() >= 12 && t.Hour() <= 18:
		return GoodAfternoon
	default:
		return GoodEvening
	}
}

func main() {
	fmt.Println(greeting(time.Now()))
}
