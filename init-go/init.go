package main

import "fmt"

var name string

func init() {
	name = "Golang cun rox bosss"
}

func main() {
	fmt.Println("Brooo this is the init function in demo ",name)
}

// The init function will only be called once despite multiple import statement
// Knowing this, we can use this feature for things like database connections
// The init function is not explicitely called, the init function is run by go 
// implicitely 

/*Order of execution of init function in go

to note - Things in go are typically initialized in the order in declaration

1 - If there are 2 files in the same package
	
	db
		- connection.go
		- query.go
	
	# if query.go depends on connection.go, connection.go init function will be run first 


*/
