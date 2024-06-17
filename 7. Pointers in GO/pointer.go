package main

import "fmt"

func main() {
	age := 24

	var agePoint *int = &age

	fmt.Println("Age Pointer:", agePoint)  // Outputs The Memory Address [0x14000112018]
	fmt.Println("Age Pointer:", *agePoint) // Outputs The Value in the Address [24]

	adultYears := getAdultYears(agePoint)
	fmt.Println("Adult Years", adultYears)
}

func getAdultYears(age *int) int {
	return *age - 18
}
