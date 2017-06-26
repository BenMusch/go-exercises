package clock

import "fmt"

const testVersion = 4
const minutesPerHour = 60
const hoursPerDay = 24

type Clock struct {
	hour, minute int
}

// pads an int with an extra string if it's only 1 character long
func pad(num int) string {
	if num < 10 {
		return fmt.Sprintf("0%v", num)
	} else {
		return fmt.Sprintf("%v", num)
	}
}

// normalizes the minutes/hours in a clock for negative values or overflows
func (clock Clock) normalize() Clock {
	// normalize minutes
	clock.hour += clock.minute / minutesPerHour
	clock.minute = clock.minute % minutesPerHour

	if clock.minute < 0 {
		clock.minute = minutesPerHour + clock.minute
		clock.hour += -1
	}

	// normalize hours
	clock.hour = clock.hour % hoursPerDay

	if clock.hour < 0 {
		clock.hour = hoursPerDay + clock.hour
	}

	return clock
}

func New(hour, minute int) Clock {
	return Clock{hour, minute}.normalize()
}

func (clock Clock) String() string {
	return fmt.Sprintf("%v:%v", pad(clock.hour), pad(clock.minute))
}

func (clock Clock) Add(minutes int) Clock {
	clock.minute += minutes
	return clock.normalize()
}
