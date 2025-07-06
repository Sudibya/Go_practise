package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello", name)
}

func sum(a int, b int) int {
	return a + b
}

func divideNumbers(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil	
}

func main() {
	greet("Sudibya")
	fmt.Println(sum(1, 2))
	fmt.Println(divideNumbers(1, 0))
}
