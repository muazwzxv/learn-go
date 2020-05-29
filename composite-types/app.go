package main

import(
	"fmt"
)

func array() {

	days := []string{"Monday", "Tuesday" ,"Wednesday", "Thursday", "Friday", "Saturday"}

	weekdays := days[0:5]
	fmt.Println(weekdays)
}

func hashMaps() {
	personSalary := map[string]float64{
		"muaz": 10000.00,
		"arep": 10000.00,
		"lano": 10000.00,
		"copia": 10000.00,
	}

	ptrPersonSalary := &personSalary

	// Printing the variable
	fmt.Println(personSalary)

	//Printing the pointer
	fmt.Println(ptrPersonSalary)

	// changing the value from the pointer
	personSalary["muaz"] = 20000

	fmt.Println(personSalary)
}

func main() {
	hashMaps()
}

