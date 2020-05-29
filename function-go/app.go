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

func main() {
	fmt.Println(concatString("Muaz", "wazir"))

	x, err := concatStringWithError("", "")
	fmt.Println(x, err)
}
