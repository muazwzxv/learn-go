package main

import(
	"fmt"
	"errors"
)

/*
To note that golang does not support function overloading
function with similar role must have different names and
cannot be overloaded
*/

// function with one return type
func concatString(first string, last string) (string) {
	return first + " " + last
}

// function with multiple return type
func concatStringWithError(first string, last string) (string, error) {
	if first == "" || last == "" {
		return " ",  errors.New("One of the strings is empty")
	}

	return first + " " + last, nil
}

func anonymousFunction() func() int {
	// Define the function to return a function
	var x int

	return func() int {
		// The function has access to x defined in the parent function
		x++

		//return x + 1
		return x
	}
}

func main() {
	fmt.Println(concatString("Muaz", "wazir"))

	x, err := concatStringWithError("", "")
	if err != nil {
		// Handler if error occurs
		fmt.Println("You got an error there boi")
		x = "Error name"
	}

	fmt.Println("Name :", x)

	myFunction := anonymousFunction()
	myFunction2 := anonymousFunction()

	// Has similar effect to static variable ?
	fmt.Println(myFunction())
	fmt.Println(myFunction())

	// Has similar effect to static variable ?
	fmt.Println(myFunction2())
	fmt.Println(myFunction2())
}






