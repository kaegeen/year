package main

import (
	"fmt"
)

// Function to check if a year is a leap year
func isLeapYear(year int) bool {
	// A leap year is divisible by 4, but not divisible by 100, unless it's also divisible by 400
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		return true
	}
	return false
}

func main() {
	var year int

	// Ask the user for the year
	fmt.Print("Enter a year: ")
	fmt.Scan(&year)

	// Calculate the number of days in the year
	daysInYear := 365
	if isLeapYear(year) {
		daysInYear = 366
	}

	// Output the result
	fmt.Printf("The year %d has %d days.\n", year, daysInYear)
}
