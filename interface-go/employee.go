package main

import(
	"fmt"
	"errors"
)

// All the functions in these interface has to be satisfied
// By all the struct implementing the interface
type Employee interface {
	GetName() string
	GetLanguageOfPreference() string
	GetAge() (int, error)
	GetExperience() (int, error)
}

// Satisfying interfaces
type Developer struct {
	Name string
	Language string
	Age int
	Experience int
}

func (e Developer) GetName() string {
	return "The dev name is " + e.Name
}

func (e Developer) GetLanguageOfPreference() string {
	return "The language of preference of this developer " + e.Language
}

func (e Developer) GetAge() (int, error) {
	if e.Age == 0 {
		return -1, errors.New("The age is not provided")
	}

	return e.Age, nil
}

func (e Developer) GetExperience() (int, error) {
	if e.Experience == 0 {
		return -1, errors.New("The year of experience is not provided")
	}

	return e.Experience, nil
}

func main() {

	var devs []Employee
	muash := Developer{"Muaz", "Golang", 21, 5,}
	afwina := Developer{"Afrina", "Java", 20, 5,}
	awep := Developer{"Arep", "Python3", 21, 5,}

	//fmt.Println(muash.GetLanguageOfPreference())
	//fmt.Println(muash.GetName())
	//fmt.Println(muash.GetAge())
	//fmt.Println(muash.GetExperience())

	devs = append(devs, muash)
	devs = append(devs, awep)
	devs = append(devs, afwina)

	for _, val := range devs {
		fmt.Println(val.GetName())
		fmt.Println(val.GetAge())
		fmt.Println(val.GetLanguageOfPreference())
		fmt.Println(val.GetExperience())
	}
}




