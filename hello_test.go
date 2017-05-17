package main

// Testing https://www.golang-book.com/books/intro/12

import (
	"fmt"
	"testing"
	"time"
)

func tenFortyFiveMorning() time.Time {
	return time.Date(2017, time.May, 11, 10, 45, 0, 0, time.UTC)
}

func sixThirtyAfternoon() time.Time {
	return time.Date(2017, time.May, 11, 18, 30, 0, 0, time.UTC)
}

func elevenFifteenNight() time.Time {
	return time.Date(2017, time.May, 11, 23, 15, 0, 0, time.UTC)
}

func TestMorningGreeting(t *testing.T) {
	//testTime := time.Date(2017, time.May, 11, 18, 30, 0, 0, time.UTC)
	text := greeting(tenFortyFiveMorning())
	expected := "Good morning"
	if text != expected {
		msg := fmt.Sprintf("Expected %s, got", expected)
		t.Error(msg, text)
	}
}

func TestAfternoonGreeting(t *testing.T) {
	//testTime := time.Date(2017, time.May, 11, 18, 30, 0, 0, time.UTC)
	text := greeting(sixThirtyAfternoon())
	expected := "Good afternoon"
	if text != expected {
		msg := fmt.Sprintf("Expected %s, got", expected)
		t.Error(msg, text)
	}
}

func TestEveningGreeting(t *testing.T) {
	text := greeting(elevenFifteenNight())
	expected := "Good evening"
	if text != expected {
		msg := fmt.Sprintf("Expected %s, got", expected)
		t.Error(msg, text)
	}
}
