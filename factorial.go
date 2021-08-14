package main

import "fmt"

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		result := value * factorialRecursive(value-1)
		return result
	}
}

func main() {
	resultFactorialRecursive := factorialRecursive(5)
	fmt.Println(resultFactorialRecursive)
}
