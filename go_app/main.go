// Created by: Lamees Hemdan
// Created on: May 2023

// this function lets you enter age and day of the week and tells you if you can get in for free or not.
package main

import (
	"fmt"
)

func main() {
	// This function lets you enter age and day of the week and tells you if you can get in for free or not.
	var age int
	var day string
	fmt.Println("Enter your age: ")
	fmt.Scanln(&age)
	fmt.Println("Enter the day of the week: ")
	fmt.Scanln(&day)

	if (day == "tuesday" || day == "thursday") || (age > 13 && age < 18) {
		fmt.Println("You can get in for free!")
	} else {
		fmt.Println("You have to pay.")

		fmt.Println("\nDone.")
	}
}
