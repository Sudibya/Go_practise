package main

import "fmt"

var name string = "Sudibya"

const pi = 3.14

const (
	StausOk = 200
	StausNotFound = 404
	StausServerError = 500
)


// Different Variable types we have in Go
// Numeric Type
var integer int = 24
var decimalNum float64 = 24.2
var complexNum complex64 = 24 + 24i


// String and Boolean Type
var str string = "Hello, World!"
var isTrue bool = true

// Zero Values
var defalutInt int
var defalyFloat float64
var defalyBool bool
var defalyString string







func main() {
	age := 24 // Short declaration, type inference
	fmt.Println(name, age, pi)
	fmt.Println(StausOk, StausNotFound, StausServerError)
}