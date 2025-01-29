package main

import "fmt"

func main() {
	age := 32

	agePointer := &age

	fmt.Println("Age:", *agePointer)
	// fmt.Println("Adult years:", getAdultYears(&age))
	getAdultYears(agePointer)
	fmt.Println("Adult years:", age)
}

func getAdultYears(age *int) {
	*age = *age - 18 // this will mutate the pointer value
}

// func getAdultYears(age *int) int {
// 	return *age - 18
// }
