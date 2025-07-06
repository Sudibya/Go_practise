package main

import "fmt"

// if/else statements
func main() {
	x := 10
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else if x == 10 {
		fmt.Println("x is equal to 10")
	} else {
		fmt.Println("x is less than 10")
	}

	// switch statements
	day := 4
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Sunday")
	}

    // For Loop

	for i:=0; i<10; i++ {
		fmt.Println(i);
	}

	// While Loop
    j:= 5
	for j <= 10{
		fmt.Println(j)
		j++
	}	

	
}