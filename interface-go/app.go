package main

import(
	"fmt"
)

// function that takes an empty interface as an arg
func myFunction(x interface{}){
	fmt.Println(x)
}

type Employee interface {
	Introduction()
}

type Programmer struct {
	Name string
	Ic string
	Age int
	Salary float64
}

type Cleaner struct {
	Name string
	Ic string
	Age int
	Salary float64
}

type Manager struct {
	Name string
	Ic string
	Age int
	Salary float64
}


// Introduction() defintions
func (x Programmer) Introduction() {
	fmt.Println("Name :", x.Name)
	fmt.Println("Ic :", x.Ic)
	fmt.Println("Age :", x.Age)
	fmt.Println("Salary :", x.Salary)
}

func (x Cleaner) Introduction() {
	fmt.Println("Name :", x.Name)
	fmt.Println("Ic :", x.Ic)
	fmt.Println("Age :", x.Age)
	fmt.Println("Salary :", x.Salary)
}

func (x Manager) Introduction() {
	fmt.Println("Name :", x.Name)
	fmt.Println("Ic :", x.Ic)
	fmt.Println("Age :", x.Age)
	fmt.Println("Salary :", x.Salary)
}

func main() {

	num := 21
	name := "Muaz Bin Wazir"
	salary := 10000.00
	// workaround for unsused variable error in golang
	_, _ = name, num

	// Any type can be pass into a fucntion that accepts an empty interface
	myFunction(salary)


	manager := Manager{"Afrina", "000000-00-0000", 20, 20000,}
	programmer := Programmer{"Muaz", "0000000-00-0000", 21, 10000,}
	cleaner := Cleaner{"Arep", "000000-00-0000", 21, 200,}

	manager.Introduction()
	programmer.Introduction()
	cleaner.Introduction()
}





