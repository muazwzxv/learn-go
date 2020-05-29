package main

import "fmt"

type person struct {
	name string
	age int
	nationality string
	ic string
	salary float64
}


func main() {
	x := person{"muaz", 21, "Malaysia", "000000-00-0000", 6000.00}
	ptr  := &x
	// Displaying the pointer to a type
	fmt.Println(ptr)

	// Go allows accessing the field of a struct using pointers without dereference
	// it explicitely
	fmt.Println("\nPrinting attributes in a pointer WITHOUT explicitely dereferencing it")
	fmt.Println(ptr.name)
	fmt.Println(ptr.age)
	fmt.Println(ptr.nationality)
	fmt.Println(ptr.ic)
	fmt.Println(ptr.salary)

	// this is how we do it if we were to explicitely dereference it
	fmt.Println("\nPrinting attributes in a pointer WITH explicitely dereferencing it")
	fmt.Println((*ptr).name)
	fmt.Println((*ptr).age)
	fmt.Println((*ptr).nationality)
	fmt.Println((*ptr).ic)
	fmt.Println((*ptr).salary)


	// updating the value of struct member
	ptr.name = "muaz semakin kacak"
	ptr.salary = 10000.00

	// Printing the pointer
	fmt.Println("\nPrinting the pointer")
	fmt.Println(ptr)

	// Printing the actual variable
	fmt.Println("\nPrinting the actual variable")
	fmt.Println(x)
}

