package main

import "fmt"

var count int

func init() {
	fmt.Println("First init executed")
	count += 5
}

func init() {
	fmt.Println("Second init executed")
	count += 10
}

func main() {
	fmt.Println("run")
	fmt.Println("Val of count :", count)
}


/* 

* When they are multiple init function in a file, order of execution is based on order of  declaration

Reason why you would need multiple init functions

1 - A way to break down one massive complicated init function
2 - Avoid having one monolith and intertwind logic in a single init

** To note that developer has to be very careful with the order of declaration to ensure proper flow of execution of the init function *** 


Conclusion - Init is very handy and a thing to use when structuring a Go project

*/
