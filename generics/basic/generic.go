package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Before ", slice)
	fmt.Println("After ", reverse(slice))
}

func reverse[T any](arr []T) []T {
	left := len(arr)
	right := make([]T, left)

	for i, element := range arr {
		right[left-i-1] = element
	}

	return right
}

/*
1, 2, 3, 4, 5

0 - 1 - 2 - 3- 4

[left-i-1] = 5 - 0 - 1 = 4
[4] = 1

[left-i-1] = 5 - 1 - 1 = 3
[3] = 2

[left-i-1] = 5 - 2 - 1 = 2
[2] = 3

[left-i-1] = 5 - 3 - 1 = 1
[1] = 4

[left-i-1] = 5 - 4 - 1 = 0
[0] = 5

*/
