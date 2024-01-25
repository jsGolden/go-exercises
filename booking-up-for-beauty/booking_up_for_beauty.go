package booking

import (
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	t, _ := time.Parse("January 2, 2006 15:04:05", date)
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	//"Thursday, July 25, 2019 13:45:00"
	t, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	// in -> "7/25/2019 13:45:00"
	//out -> Thursday, July 25, 2019, at 13:45
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	formattedTime := t.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
	return formattedTime
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	//t := time.Date(1995,time.September,22,13,0,0,0,time.UTC)
	t := time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
	return t
}
