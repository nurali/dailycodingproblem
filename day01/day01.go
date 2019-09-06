package main

import "fmt"

// Given a list of numbers and a number k, return whether any two numbers from the list add up to k.
// For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

func main() {
	numbers := []int{10, 15, 3, 7}
	no := 17
	fmt.Println(CanSumUp(numbers, no))
}

func CanSumUp(numbers []int, no int) bool {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if j == i {
				continue
			}
			if numbers[i]+numbers[j] == no {
				return true
			}
		}
	}
	return false
}
