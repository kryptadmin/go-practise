package main

import "fmt"

func main() {
	var favoriteColor = "red"
	fmt.Println("Favourite color is ", favoriteColor)

	birthYear, ageInYears := 1985, 37
	fmt.Println("My Birth year is ", birthYear, "My age is ", ageInYears)

	var (
		firstInitial = "D"
		secondInitial = "D"
	)

	fmt.Println("My Initials are ", firstInitial, secondInitial)

	var ageInDays int
	ageInDays = 365 * ageInYears
	fmt.Println("My age in days is ", ageInDays)
}