package main

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHello(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hello Suite")
}

var _ = Describe("Hello", func() {
	var (
		morningTime   time.Time
		afternoonTime time.Time
		eveningTime   time.Time
	)

	BeforeEach(func() {
		morningTime = time.Date(2017, time.May, 11, 10, 45, 0, 0, time.UTC)
		afternoonTime = time.Date(2017, time.May, 11, 18, 30, 0, 0, time.UTC)
		eveningTime = time.Date(2017, time.May, 11, 23, 15, 0, 0, time.UTC)
	})

	Describe("Greeting method", func() {
		Describe("In the morning", func() {
			It("returns Good morning", func() {
				Expect(greeting(morningTime)).To(Equal("Good morning"))
			})
		})
		Describe("In the afternoon", func() {
			It("returns Good afternoon", func() {
				Expect(greeting(afternoonTime)).To(Equal("Good afternoon"))
			})
		})
		Describe("In the evening", func() {
			It("returns Good evening", func() {
				Expect(greeting(eveningTime)).To(Equal("Good evening"))
			})
		})
	})
})
